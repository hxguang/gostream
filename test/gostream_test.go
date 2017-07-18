package test

import (
	"testing"
	"gostream"
	"fmt"
	"time"
)

func TestFilter(t *testing.T){
	t1 := time.Now()
	orgin := []int{1,2,3,4,5,6}
	result := gostream.Filter(orgin,func(item interface{}) bool {
		return item.(int)%2 == 0
	})
	duration := time.Since(t1)
	fmt.Println("gostream.Filter execute duration:",duration)
	fmt.Println(result)
}

func TestMap(t *testing.T){
	t1 := time.Now()
	orgin := []int{1,2,3,4,5,6}
	result := gostream.Map(orgin,func(item interface{})interface{}{
		i := item.(int)
		return i*2
	})
	duration := time.Since(t1)
	fmt.Println("gostream.Map execute duration:",duration)
	fmt.Println(result)
}

func TestReduce(t *testing.T){
	t1 := time.Now()
	sliceInt := make([]int,1000)
	for i := range sliceInt{
		sliceInt[i] = int(i) + int(1)
	}
	result := gostream.Reduce(sliceInt,doAdd,int(10))
	t2 := time.Since(t1)
	fmt.Println("gostream.Reduce execute duration:",t2)
	fmt.Println(result)
}

func doAdd(first interface{},second interface{}) interface{} {
	f := first.(int)
	s := second.(int)
	return f + s
}
