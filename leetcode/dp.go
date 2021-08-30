// 动态规划
package leetcode

import (
	"fmt"
	"github.com/iamxuwenjin/blog/model"
)

func step(num int) int {
	dp := make([]int, 0, num)
	for i := 1; i <= num; i++ {
		if i == 1 {
			dp = append(dp, 1)
		} else if i == 2 {
			dp = append(dp, 2)
		} else {
			dp = append(dp, dp[i-2]+dp[i-3])
		}

	}
	return dp[num-1]
}

func robot(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
			fmt.Println(dp)
		}
	}
	return dp[m-1][n-1]
}

//64. 最小路径和
func minPathSum(grid [][]int) int {
	var dp = grid
	rows, columns := len(grid), len(grid[0])
	for i, num1 := range grid {
		for j, _ := range num1 {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 && j > 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 && i > 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + model.Min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[rows-1][columns-1]
}

//72. 编辑距离
//func longestPalindrome(s string) string {
//	n := len(s)
//	dp := make([]int, n)
//	for i := 0; i < n; i++ {
//		if i < 2 {
//			if i == 1 && s[i] == s[i-1] {
//				dp[i] = 2
//			} else {
//				dp[i] = 1
//			}
//		} else {
//			if i-dp[i-1]-1 < 0 {
//				if s[i] == s[i-1] {
//					dp[i] = dp[i-1] + 1
//				} else {
//					dp[i] = 1
//				}
//				continue
//			}
//			preStr := s[i-dp[i-1]-1]
//			if s[i] == preStr {
//				dp[i] = dp[i-1] + 2
//			} else if s[i] == s[i-1] {
//				dp[i] = 2
//			} else {
//				dp[i] = 1
//			}
//		}
//	}
//	maxNum := 0
//	maxIdx := 0
//	for idx, num := range dp {
//		if num > maxNum {
//			maxNum = num
//			maxIdx = idx
//		} else {
//			continue
//		}
//	}
//	return s[maxIdx-maxNum+1 : maxIdx+1]
//}

//买卖股票的最佳时机
func maxProfit(prices []int) int {
	minPrices := prices[0]
	ans := 0
	for _, num := range prices {
		if num < minPrices {
			minPrices = num
		}
		fmt.Println(num - minPrices)
		//if ans < num - minPrices{
		//	ans = num - minPrices
		//}
		ans = model.Max(ans, num-minPrices)
	}
	return ans
}

//三角形最小路径和
func minimumTotal(triangle [][]int) int {
	var dp = triangle
	for i, numList := range triangle {
		for j, _ := range numList {
			if i == 0 {
				dp[i][j] = triangle[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = model.Min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	ans := triangle[len(triangle)-1][0]
	for _, num := range triangle[len(triangle)-1] {
		if num < ans {
			ans = num
		}
	}
	return ans
}

//53. 最大子序和
func maxSubArray(nums []int) int {
	var dp = nums
	ans := dp[0]
	for idx, _ := range nums {
		if idx == 0 {
			dp[idx] = nums[idx]
		} else if dp[idx-1] <= 0 {
			dp[idx] = nums[idx]
		} else {
			dp[idx] = dp[idx-1] + nums[idx]
		}
		if ans < dp[idx] {
			ans = dp[idx]
		}
	}
	return ans
}

//198. 打家劫舍
func rob(nums []int) int {
	var dp = nums
	ans := 0
	for idx, _ := range nums {
		if idx <= 1 {
			dp[idx] = nums[idx]
		} else if idx == 2 {
			dp[idx] = dp[idx] + dp[idx-2]
		} else {
			dp[idx] = model.Max(dp[idx-2], dp[idx-3]) + nums[idx]
		}
		ans = model.Max(dp[idx], ans)
	}
	return ans
}

//5. 最长回文子串
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLen := 1
	startIdx := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			if i == j {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}
	// 枚举所有字符串
	for i := 2; i <= len(s); i++ {
		for j := 0; j < len(s); j++ {
			r := j + i - 1
			if r >= len(s) {
				continue
			}
			if s[j] != s[r] {
				continue
			} else {
				if r-j < 3 {
					dp[j][r] = true
				} else {
					dp[j][r] = dp[j+1][r-1]
				}
			}
			if dp[j][r] && r-j+1 > maxLen {
				maxLen = r - j + 1
				startIdx = j
			}
		}
	}
	return s[startIdx : startIdx+maxLen]
}

//回文子串
func countSubstrings(s string) int {
	ans := len(s)
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			if j == i {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}
	for i := 2; i <= len(s); i++ {
		for j := 0; j < len(s); j++ {
			r := i + j - 1
			if r >= len(s) {
				continue
			}
			if s[j] != s[r] {
				continue
			} else {
				if r-j < 3 {
					dp[j][r] = true
					ans++
				} else {
					if dp[j+1][r-1] {
						dp[j][r] = dp[j+1][r-1]
						ans++
					}
				}
			}
		}
	}
	return ans
}

//122. 买卖股票的最佳时机 II
func maxProfit2(prices []int) int {
	maxP := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			maxP += prices[i] - prices[i-1]
		}
	}
	return maxP
}

// 最长公共子串
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	ans := 0
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if ans < dp[i+1][j+1] {
					ans = dp[i+1][j+1]
				}
			}
		}
	}
	return ans
}

//303. 数组不可变
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for idx, val := range nums {
		sums[idx+1] = sums[idx] + val
	}
	return NumArray{
		nums: sums,
	}
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.nums[right+1] - na.nums[left]
}

//560.和为k的子数组
func subArraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

//和等于 k 的最长子数组长度
func subArrayLEN(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count = model.Max(i-m[pre], count)
		}
		if _, ok := m[pre]; !ok {
			m[pre] = i
		}
	}
	return count
}
