package main

import (
	"log"
	"sort"
)

func main() {
	arrs := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(arrs)
	log.Println(res)
	return
}
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -1 * nums[first]
		third := n - 1
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second >= third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

// -1,0,1,2,-1,-4
func threeSum2(nums []int) [][]int {
	length := len(nums)
	res := make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < length; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -1 * nums[first]
		third := length - 1
		for second := first + 1; second < length; second++ {
			if second > first && nums[second] == nums[second]-1 {
				continue
			}

			if second < third && nums[second]+nums[third] > target {
				third--
			}
			if second >= third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
