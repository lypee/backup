package main

import "log"

func main() {
	nums := []int{1, 2, 3, 4, 6, 7, 8}
	sort(nums)
}

func sort(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]%2 == 1 {
			l++
		}
		for l < r && nums[r]%2 != 1 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	log.Println(nums)
}
