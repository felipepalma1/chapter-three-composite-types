package chapter_three_composite_types

import (
	"github.com/felipepalma1/chapter-three-special-array/operations"
)

func SimpleArray() ([3]int, error) {
	var x [3]int

	if x[0] == 0 {
		return [3]int{}, nil
	}
	return x, nil
}

func SimpleInitializedArray() [3]int {
	var x = [3]int{10, 20, 30}
	return x
}

func NewSpecialArray() [12]int {
	return operations.NewSparseArray()
}

func LoadingASlice() []int {
	var x []int
	x = append(x, 10)
	x = append(x, 10)
	x = append(x, 10)
	x = append(x, 10)
	x = append(x, 10, 20, 30, 40)

	return x
}

func MakingSlice() []int {
	x := make([]int, 1, 20)

	for range x {
		x = append(x, 100)
	}

	return x
}
