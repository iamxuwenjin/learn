package model

import "math"

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// εθδΊε₯
func round(x float64) int {
	return int(math.Floor(x + 0/5))
}
