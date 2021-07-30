package leetcode

func factorial(num int) int {
	if num == 1 {
		return num
	}
	return num * factorial(num-1)
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	} else {
		return 1 / quickMul(x, -n)
	}
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	// 不考虑复杂度
	//return x * myPow(x, n-1)
	y := myPow(x, n/2)
	mod := n % 2
	if mod == 1 {
		return y * y * x
	} else {
		return y * y
	}
}

// 快速排序
func quickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, start, end int) []int {
	if start < end {
		midIndex := partition(arr, start, end)
		_quickSort(arr[start:midIndex], start, midIndex-1)
		_quickSort(arr[midIndex:end], midIndex+1, end)
	}
	return arr
}

//[5,4,6,7,3,2]
func partition(arr []int, start, end int) int {
	pivot := arr[start]
	midIdx := start + 1
	for idx := pivot + 1; idx <= end; idx++ {
		if arr[idx] < pivot {
			swap(arr, idx, midIdx)
			midIdx++
		}
	}
	return midIdx - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
