package leetcode

import (
	"github.com/iamxuwenjin/blog/model"
	"math"
	"strconv"
)

//415. 字符串相加
func addStrings(num1 string, num2 string) string {
	ans := ""
	add := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			t := num1[i] - '0'
			x = int(t)
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

// 归并排序
func mergeSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	// 切分
	mid := n / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// 合并
	res := merge(left, right)
	return res
}

// 组合两个有序数组并且排序
func merge(num1 []int, num2 []int) []int {

	target := make([]int, 0, len(num1)+len(num2))
	for i, j := 0, 0; i <= len(num1)-1 || j <= len(num2)-1; {
		// 当一个列表到达最后一个元素时，直接添加另一个列表剩余元素到target
		if i > len(num1)-1 {
			target = append(target, num2[j:]...)
			break
		}
		if j > len(num2)-1 {
			target = append(target, num1[i:]...)
			break
		}
		//	取出列表当前值进行比较，取出min(a, b)放入target
		if num1[i] < num2[j] {
			target = append(target, num1[i])
			i++
		} else {
			target = append(target, num2[j])
			j++
		}
	}
	return target
}

//剑指 Offer 51. 数组中的逆序对
func reversePairs(nums []int) int {
	return findReversePairs(nums, 0, len(nums)-1)
}

func findReversePairs(nums []int, start, end int) int {
	// 递归结束条件
	if start >= end {
		return 0
	}
	// 分割数组
	mid := start + (end-start)/2
	cnt := findReversePairs(nums, start, mid) + findReversePairs(nums, mid+1, end)
	// 合并数组
	tmp := make([]int, 0, 0)
	i, j := start, mid+1
	// 当其中一个数据遍历至-1时，结束循环
	for i <= mid && j <= end {
		if nums[i] < nums[j] {
			tmp = append(tmp, nums[i])
			cnt = j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}

	}
	//当左边列表未遍历完整时
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	//当右边列表未遍历完整时
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	//数组替换
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}

// 41. 缺失的第一个正数
func firstMissingPositive(nums []int) int {
	// hash实现
	//numsMap := make(map[int]int)
	//for _, i := range nums{
	//	numsMap[i] = i
	//}
	//n := len(nums)
	//i := 1
	//for ; i < n; i ++{
	//	if _, ok := numsMap[i]; ok == false{
	//		return i
	//	}
	//}
	//return i + 1
	// 原地哈希
	// 答案只能出现在 1 - n+1中，n = len(nums), 列表索引恰好能代表这个范围
	n := len(nums)
	for _, num := range nums {
		if num > n || num < 1 {
			continue
		} else {
			nums[num-1] = int(math.Abs(float64(nums[num-1])))
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < n-1; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	max := 0
	// 查找最大堆的香蕉个数
	for _, num := range piles {
		if num > max {
			max = num
		}
	}
	if h == n {
		//问题转换成求列表中最大的数
		return max
	} else {
		i, j := 0, max
		for i < j {
			mid := (i + j) / 2
			ok := possible(mid, h, piles)
			if ok {
				j = mid
			} else {
				i = mid + 1
			}
		}
		return i
	}
}

func possible(k, h int, piles []int) bool {
	total := 0
	for _, num := range piles {
		total += num / k
		if num%k > 0 {
			total++
		}
	}
	return total <= h
}

// 160. 相交链表
func getIntersectionNode(headA, headB *model.MyLinkedList) *model.Node {
	nodeMap := make(map[*model.Node]bool)
	tempA := headA.Head
	tempB := headB.Head
	for temp := tempA; temp != nil; temp = temp.Next {
		nodeMap[temp] = true
	}
	for temp := tempB; temp != nil; temp = temp.Next {
		if _, ok := nodeMap[temp]; ok {
			return temp
		}
	}
	return nil
}

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	diff := nums[0]
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2
		//判断区间是否是局部有序的
		if nums[mid]-diff-mid >= 0 {
			//l部分确定有序，r部分仍然可能是首尾有序
			//判断目标值范围
			if nums[l] <= target && target < nums[mid] {
				// 位于固定有序区
				//if target > nums[mid]{
				//	l = mid + 1
				//}
				r = mid - 1
			} else if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				return mid
			}
		} else {
			// l部分仍然可能是首尾有序，R部分确定有序
			// 确定目标范围
			if target < nums[mid] || target > nums[r] {
				//
				r = mid - 1
			} else if target > nums[mid] && target <= nums[r] {
				//
				l = mid + 1
			} else {
				return mid
			}
		}

	}
	return -1
}

// 11. 盛水最多的容器
func maxArea(height []int) int {
	// 计算公式 min(a, b) * (index(b) - index(a))
	ans := 0
	for l, r := 0, len(height)-1; l <= r; {
		var area int
		if height[l] < height[r] {
			area = height[l] * (r - l)
			l++
		} else {
			area = height[r] * (r - l)
			r--
		}
		if area > ans {
			ans = area
		}
	}
	return ans
}

//42. 接雨水
func trap(height []int) int {
	ans := 0
	leftMax := 0
	rightMax := 0
	for l, r := 0, len(height)-1; l <= r; {
		if height[l] < height[r] {
			if height[l] > leftMax {
				leftMax = height[l]
			}
			ans += leftMax - height[l]
			l++
		} else {
			if height[r] > rightMax {
				rightMax = height[r]
			}
			ans += rightMax - height[r]
			r--
		}
	}
	return ans
}

//303. 最长递增子序列DP
func lengthOfLIS(nums []int) int {
	ansList := make([]int, 0, len(nums))
	for _, nums := range nums {
		if len(ansList) == 0 {
			ansList = append(ansList, nums)
			continue
		}
		// 判断目标队列最后一个元素与当前num的大小
		if ansList[len(ansList)-1] > nums {
			ansList[len(ansList)-1] = nums
		} else {
			ansList = append(ansList, nums)
		}
	}
	return len(ansList)
}

//1143. 最长公共子序列DP
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1) // m+1 行
	for i := range dp {
		dp[i] = make([]int, n+1) // n+1列
	}
	for i, str1 := range text1 {
		for j, str2 := range text2 {
			if str1 == str2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = model.Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

//136. 只出现一次的数字
func singleNumber(nums []int) int {
	ansMap := make(map[int]bool)
	var ans int
	for _, nums := range nums {
		if ansMap[nums] {
			delete(ansMap, nums)

		} else {
			ansMap[nums] = true
		}
	}
	for idx, val := range ansMap {
		if val {
			ans = idx
		}
	}
	return ans
}

//131. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	// 找起点
	start := 0
	for ; start < n-1; start++ {
		if gas[start] >= cost[start] {
			break
		}
	}
	run := start
	total := gas[start]
	var pre int
	for {
		run++
		if run == n {
			run = 0
		}
		if run == 0 {
			pre = n - 1
		} else {
			pre = run - 1
		}
		total -= cost[pre]
		if total < 0 {
			run = pre
			break
		}
		total += gas[run]
		if run == start {
			break
		}
	}
	if run == start {
		return start
	} else {
		return -1
	}
}
