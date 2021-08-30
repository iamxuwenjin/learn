package leetcode

import (
	"fmt"
	"github.com/iamxuwenjin/blog/model"
	"strconv"
	"testing"
)

func TestAddString(t *testing.T) {
	a := "2385423"
	b := "123143"
	c := "1sdfasd"
	length := len(c)
	for i := 0; i < length; i++ {
		fmt.Printf("%c", i)
	}
	for _, i := range c {
		fmt.Printf("%c", i)
	}
	res := addStrings(a, b)
	intA, _ := strconv.Atoi(a)
	intB, _ := strconv.Atoi(b)
	sums := strconv.Itoa(intA + intB)
	fmt.Println(sums == res)
}

func TestMerge(t *testing.T) {
	a := []int{1, 4, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10, 11, 34}
	res := merge(a, b)
	fmt.Println(res)
}

func TestMergeSort(t *testing.T) {
	nums := []int{1, 4, 3, 11, 5, 3, 2}
	res := mergeSort(nums)
	fmt.Println(res)

}

// 向零取整
func TestDev(t *testing.T) {
	fmt.Println(5 / 2)
	fmt.Printf("%f", 5.0/2.0)
}

func TestReversePairs(t *testing.T) {
	nums := []int{1, 8, 4, 5, 2, 11, 9}
	reversePairs(nums)
}

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{2, 3, 4, 5, 6, 7}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}

func TestPossible(t *testing.T) {
	nums := []int{5, 10, 20}
	res := minEatingSpeed(nums, 4)
	//fmt.Println(int(math.Ceil(2.0 / 6.0)))
	fmt.Println(res)
}

func TestGetIntersectionNode(t *testing.T) {
	num1 := []int{1, 4, 5, 6, 7, 9}
	num2 := []int{2, 54, 4, 6, 7, 9}
	nodeList1 := model.GetListNode(num1)
	nodeList2 := model.GetListNode(num2)
	res := getIntersectionNode(nodeList1, nodeList2)
	fmt.Println(*res)
}

func TestSearch(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums := []int{4, 2, 3}
	res := search(nums, 8)
	fmt.Println(res)
}

func TestMaxArea(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//nums := []int{4, 2, 3}
	res := maxArea(nums)
	fmt.Println(res)
}

func TestTrap(t *testing.T) {
	nums := []int{4, 2, 0, 3, 2, 5}
	res := trap(nums)
	fmt.Println(res)
}

func TestLenOfList(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2, 3}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}

func TestLongestCommonSubsequence(t *testing.T) {
	str1 := "foolish"
	str2 := "fool"
	res := longestCommonSubsequence(str1, str2)
	fmt.Println(res)
}

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 1, 4, 6, 6}
	res := singleNumber(nums)
	fmt.Println(res)
}

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	res := canCompleteCircuit(gas, cost)
	fmt.Println(res)
}
