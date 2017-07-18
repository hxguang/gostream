package gostream

import (
	"reflect"
	"errors"
)

type PreCond func(elem interface{}) bool
type PreProc func(elem interface{}) interface{}
type Reducer func(lastElem interface{}, newElem interface{}) interface{}

func Filter(orgin interface{}, cond PreCond) []interface{} {
	if !isSupportedType(orgin) {
		panic(errors.New("type is not supported"))
	}
	var result []interface{}
	orginValue := reflect.ValueOf(orgin)
	len := orginValue.Len()
	var elem interface{}
	for i := 0; i < len; i++ {
		elem = orginValue.Index(i).Interface()
		if cond(elem) {
			result = append(result, elem)
		}
	}
	return result
}

func Map(orgin interface{}, proc PreProc) []interface{} {
	if !isSupportedType(orgin) {
		panic(errors.New("type is not supported"))
	}
	var result []interface{}
	orginValue := reflect.ValueOf(orgin)
	len := orginValue.Len()
	for i := 0; i < len; i++ {
		result = append(result, proc(orginValue.Index(i).Interface()))
	}
	return result
}

func Reduce(orgin interface{}, reducer Reducer, beginItem interface{}) interface{} {
	if !isSupportedType(orgin) {
		panic(errors.New("type is not supported"))
	}
	seq := wrapToSeq(orgin)
	if beginItem != nil {
		seq = append(seq, beginItem)
	}
	length := len(seq)
	if length < 2 {
		panic(errors.New("sum of elememts must be >= 2"))
	}
	first := seq[0]
	last := seq[1]
	last = reducer(first,last)
	for i := 2;i < length;i++{
		last = reducer(seq[i],last)
	}
	return last
}

func wrapToSeq(orgin interface{}) []interface{} {
	var seq []interface{}
	orginValueType := reflect.ValueOf(orgin)
	len := orginValueType.Len()
	for i := 0; i < len; i++ {
		seq = append(seq, orginValueType.Index(i).Interface())
	}
	return seq
}

func isSupportedType(orgin interface{}) bool {
	isSupported := false

	switch reflect.ValueOf(orgin).Kind() {
	case reflect.Slice:
		isSupported = true
	case reflect.Array:
		isSupported = true
	case reflect.String:
		isSupported = true
	}
	return isSupported
}
