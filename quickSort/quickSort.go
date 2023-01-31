package main

import "log"

func main() {
	nums := []int{1}
	log.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	resNums := sort(nums, 0, len(nums)-1)
	return resNums
}

func sort(nums []int, left, right int) []int {
	if left > right {
		return nil
	}
	l, r, sNum := left, right, nums[left]
	//log.Printf("l :[%d] , r :[%d] , sNum:[%+v]", l, r, nums)
	for l < r {
		for l < r && nums[r] >= sNum {
			r--
		}
		for l < r && nums[l] <= sNum {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[l], nums[left] = nums[left], nums[l]
	sort(nums, left, l-1)
	sort(nums, l+1, right)
	return nums
}
