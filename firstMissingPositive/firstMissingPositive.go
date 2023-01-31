package main

import (
	"log"
	"sort"
)

func main() {
	nums := []int{0, 1}
	log.Println(firstMissingPositive(nums))
}
func firstMissingPositive(nums []int) int {
	if len(nums) == 1 {
		if nums[0] == 1 {
			return 2
		} else {
			return 1
		}
	}
	sort.Ints(nums)
	var i int
	var moveNums int
	for i = 1; i < nums[len(nums)-1]; {
		for idx, num := range nums {
			if num < 1 {
				continue
			}
			moveNums++
			if idx > 0 && nums[idx] == nums[idx-1] {
				continue
			}
			if num == i {
				i++
				continue
			} else {
				return i
			}
		}
	}
	if moveNums == 0 {
		if nums[len(nums) -1] == 1 {
			return 2
		}
		return 1
	}
	if i == nums[len(nums)-1]+1 {
		return i
	}
	return -1
}
