package main

import "log"

func main() {
	s := "abaccdeff"
	log.Println(string(firstUniqChar(s)))
}

func firstUniqChar(s string) byte {
	countSlice := make([]int, 26)
	for _, c := range s {
		countSlice[c-'a']++
	}

	for _, c := range s {
		if countSlice[c-'a'] == 1 {
			return byte(c)
		}
	}
	return ' '
}
