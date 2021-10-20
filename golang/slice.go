package golang

import (
	"errors"
	"fmt"
)

func sliceDemo() {
	sl := make([]int, 0, 10)
	var appendFunc = func(s []int) {
		s = append(s, 10, 20, 30)
		fmt.Println(s)
	}
	fmt.Println(sl)
	appendFunc(sl)
	fmt.Println(sl)
	fmt.Println(sl[:10])
	errors.New("22")
}
