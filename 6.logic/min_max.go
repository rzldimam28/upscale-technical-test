package logic

import "errors"

var (
	ErrNotValidArr = errors.New("invalid array input")
)

func MinMax(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return []int{}, ErrNotValidArr
	}

	min := arr[0]
	max := arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return []int{min, max}, nil
}