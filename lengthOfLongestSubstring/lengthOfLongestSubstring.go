package main

import (
	"log"
	"math"
)

func main() {
	str := "pwwkew"
	log.Println(lengthOfLongestSubstring(str))
}

func max(arr ...int) int {
	res := math.MinInt64
	for _, num := range arr {
		if num > res {
			res = num
		}
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	res, left := 0, 0
	mmap := make(map[int32]int, 0)
	for idx, c := range s {
		if leftIdx, exist := mmap[c]; exist {
			left = max(left, leftIdx+1)
		}
		res = max(res, idx -left +1)
		mmap[c] = idx
	}
	return res
}


func lengthOfLongestSubstring3(s string) int{
	length := len(s)
	dp := make([][]int , length +1 )
	d
}