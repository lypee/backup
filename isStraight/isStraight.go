package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	nums := []int{0, 0, 8, 5, 4}
	log.Println(isStraight(nums))
}

func isStraight(nums []int) bool {
	sort.Ints(nums)
	jokerIdx := 0
	for idx, num := range nums {
		if num == 0 {
			jokerIdx++
			continue
		}
		if idx > 0 && nums[idx] <= nums[idx-1] {
			return false
		}
	}
	log.Println(jokerIdx)
	log.Println(nums)
	return (nums[4] - nums[jokerIdx]) <= 5
}

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

func isStraight2(nums []int) bool {
	sort.Ints(nums)
	jokerIdx := 0
	for idx, num := range nums {
		if num == 0 {
			jokerIdx++
			continue
		}
		if idx > 0 && nums[idx] > nums[idx-1] {
			return false
		}
	}
	return (nums[4] - nums[jokerIdx]) < 5
}
