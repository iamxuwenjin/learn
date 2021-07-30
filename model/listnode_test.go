package model

import (
	"fmt"
	"testing"
)

func TestGenListNode(t *testing.T) {
	nums := []int{1, 3, 54, 6}
	res := GetListNode(nums)
	fmt.Println(*res)
}

func TestConstructor(t *testing.T) {
	l := Constructor()
	//l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	res := l.Get(1)
	fmt.Println(res)
	l.DeleteAtIndex(1)
	res = l.Get(1)
	fmt.Println(res)
}

func TestAB(t *testing.T) {
	a, b := 1, 4
	fmt.Println(a, b)
}
