package main

import "log"

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	combinationSum(candidates, target)
	log.Println(res)
}

var (
	res  [][]int
	cur  []int
	nums []int
)

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	cur = make([]int, 0)
	nums = candidates
	help(target, 0)

	return res
}

func help(target, idx int) {
	if idx >= len(nums) {
		return
	}
	if target == 0 {
		res = append(res, append([]int{}, cur...))
		return
	}

	help(target, idx+1)

	if target-nums[idx] >= 0 {
		cur = append(cur, nums[idx])
		help(target-nums[idx], idx)
		cur = cur[:len(cur)-1]
	}

	return
}
