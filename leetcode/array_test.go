package leetcode

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	nums := []int{-1, -1, -1}
	//nums := []int{2, 1, -1}
	res := pivotIndex(nums)
	fmt.Println(res)
}
