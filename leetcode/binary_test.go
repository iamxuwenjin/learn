package leetcode

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 11
	fmt.Println(binarySearch(nums, target))
}
