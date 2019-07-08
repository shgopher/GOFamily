package main

import (
	"fmt"
	"testing"
)

var(
	sequeue = NewQueueSliceSequence(4)
	cycle = NewQueueSliceCycle(4)
)
func TestQueueSliceSequence_In(t *testing.T) {
	sequeue.In("a")
	sequeue.In("b")
	sequeue.In("c")
	sequeue.In("d")

}

func TestQueueSliceSequence_Out(t *testing.T) {
	fmt.Println(sequeue.Out())
	fmt.Println(sequeue.Out())
	fmt.Println(sequeue.Out())
	fmt.Println(sequeue.Out())
}

func TestNewQueueSliceSequence(t *testing.T) {
	fmt.Println("测试大大大")
	fmt.Println(sequeue.length)
	sequeue.In("大大大")
	sequeue.In("da")
	sequeue.In("dada")
	sequeue.In("dadada")
	fmt.Println(sequeue.Out())
	fmt.Println(sequeue.Out())
	fmt.Println(sequeue.Out())
	fmt.Println("///",len(sequeue.value))
}

func TestQueueSliceCycle_In(t *testing.T) {
	cycle.In("I")
	cycle.In("II")
	cycle.In("III")
	cycle.In("IV")
	cycle.In("V")
	fmt.Println("测试cycle",cycle.length)
}

func TestQueueSliceCycle_Out(t *testing.T) {
	fmt.Println(cycle.Out())
	fmt.Println(cycle.Out())
	fmt.Println(cycle.Out())
	fmt.Println(cycle.Out())
	fmt.Println(cycle.Out())

}

func TestNewQueueSliceCycle(t *testing.T) {
	cycle.In(".")
	cycle.In("..")
	cycle.In("...")
	cycle.In("....")
	cycle.Out()
	cycle.Out()
	cycle.In("xx")
	cycle.In("xx")
	fmt.Println(cycle.Out())
	fmt.Println(cycle.Out())
	fmt.Println(cycle.Out())
	fmt.Println(cycle.Out())
}