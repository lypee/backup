package main

import "log"

func main() {
	log.Println(lengthOfLongestSubstring("a"))
	log.Println(lengthOfLongestSubstring("ay"))
	log.Println(lengthOfLongestSubstring("  "))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	mmap := make(map[int32]int, 0)
	n := 0
	left := 0
	for right, c := range s {
		if idx, exist := mmap[c]; exist {
			left = max(left, idx+1)
		}
		n = max(n, right-left+1)
		mmap[c] = right
	}

	return n
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
