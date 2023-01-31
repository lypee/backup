package main

import "log"

func main() {
	nums1 := []int{}
	nums2 := []int{1}
	log.Println(merge(nums1, nums2))
}

func merge(nums1 []int, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	resNum := make([]int, 0, len1+len2)
	l, r := 0, 0
	for l < len1 && r < len2 {
		for l < len1 && r < len2 && nums1[l] <= nums2[r] {
			resNum = append(resNum, nums1[l])
			l++
		}
		for l < len1 && r < len2 && nums2[r] <= nums1[l] {
			resNum = append(resNum, nums2[r])
			r++
		}
	}
	if l < len1 {
		resNum = append(resNum, nums1[l:]...)
	}
	if r < len2 {
		resNum = append(resNum, nums2[r:]...)
	}
	return resNum
}
