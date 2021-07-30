package leetcode

import (
	"fmt"
	"testing"
)

func TestStep(t *testing.T) {
	res := step(6)
	fmt.Println(res)
}

func TestRobot(t *testing.T) {
	res := robot(7, 3)
	fmt.Println(res)
}

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	res := minPathSum(grid)
	fmt.Println(res)
}

func TestLongestPalindrome(t *testing.T) {
	str := "abbcccbbbcaaccbababcbcabca"
	res := longestPalindrome(str)
	fmt.Println(res)
}

func TestMaxProfit(t *testing.T) {
	price := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(price)
	fmt.Println(res)
}

func TestMinimumTotal(t *testing.T) {
	grid := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res := minimumTotal(grid)
	fmt.Println(res)
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	fmt.Println(res)
}

func TestRob(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	res := rob(nums)
	fmt.Println(res)
}

func TestLongest(t *testing.T) {
	str := "ab"
	res := longestPalindrome(str)
	fmt.Println(res)
}

func TestCountSubString(t *testing.T) {
	str := "aaaaa"
	res := countSubstrings(str)
	fmt.Println(res)
}

func TestProfit2(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit2(nums)
	fmt.Println(res)
}

func TestFindLength(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	res := findLength(nums1, nums2)
	fmt.Println(res)
}

func TestSubArrayLEN(t *testing.T) {
	nums2 := []int{1, -1, 5, -2, 3}
	res := subArrayLEN(nums2, 3)
	fmt.Println(res)
}
