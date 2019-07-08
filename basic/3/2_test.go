package main

import (
	"fmt"
	"testing"
)

var(
	testLink = NewQueueLinkedList()
)
func TestQueueLinkedList_In(t *testing.T) {
	testLink.In(1)
	testLink.In(2)
	testLink.In(3)
	testLink.In(4)
	testLink.In(5)
	testLink.In(6)
	testLink.In(7)
	testLink.In(8)
	fmt.Println(testLink.length)

}

func TestQueueLinkedList_Out(t *testing.T) {
	fmt.Println(testLink.Out())
	fmt.Println(testLink.Out())
	fmt.Println(testLink.Out())
	fmt.Println(testLink.Out())
	fmt.Println(testLink.Out())
	fmt.Println(testLink.Out())
	fmt.Println(testLink.Out())
	fmt.Println(testLink.Out())
	fmt.Println(testLink.length)
	fmt.Println(testLink.Out())

}
