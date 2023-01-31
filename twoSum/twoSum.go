package main

import (
	"log"
	"time"
)

func main() {
	log.Println(time.Now().Unix())
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	if len(nums) < 2 {
		return res
	}
	resMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if val, ok := resMap[diff]; ok {
			res[0] = i
			res[1] = val
			return res
		}
		resMap[nums[i]] = i
	}
	return res
}
