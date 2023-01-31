package utils

import "math"

func Max(arr ...int) int {
	res := math.MinInt64
	for i := range arr {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}

func Min(arr ...int) int {
	res := math.MaxInt64
	for i := range arr {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}
