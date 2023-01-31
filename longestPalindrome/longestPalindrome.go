package main

import "log"

func main() {
	s := "babad"
	log.Println(longestPalindrome(s))
}

//func longestPalindrome(s string) string {
//	start, end := 0, 0
//	for i := 0; i < len(s)-1; i++ {
//		s1, e1 := expend(s, i, i)
//		s2, e2 := expend(s, i, i+1)
//		if end-start < e1-s1 {
//			start, end = s1, e1
//		}
//		if end-start < e2-s2 {
//			start, end = s2, e2
//		}
//	}
//
//	return s[start : end+1]
//}

//
//func expend(s string, l, r int) (int, int) {
//
//	for l >= 0 && r < len(s) && s[l] == s[r] {
//		l--
//		r++
//	}
//	return l + 1, r - 1
//}

func longestPalindrome(s string) string {
	start, end := 0, 0
	for idx := range s {
		s1, e1 := expend(s, idx, idx)
		s2, e2 := expend(s, idx, idx+1)
		log.Printf("s1 [%d] , e1 [%d] , s2 [%d]  ,e2  [%d] ", s1, e1, s2, e2)
		if end-start < e1-s1 {
			log.Println("change ")
			start, end = s1, e1
		}
		if end-start < e2-s2 {
			log.Println("change ")
			start, end = s2, e2
		}
	}

	log.Printf("start [%d] , end  [%d] ,", start, end)

	return s[start : end+1]
}

func expend(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}
