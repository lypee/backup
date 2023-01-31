package main

import "log"

func main() {
	l4 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  2,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}

	log.Println(isPalindrome(l1))
}

//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	xx := x
//	res := 0
//	for x != 0 {
//		n := x % 10
//		res = res *10 + n
//		x /= 10
//	}
//	return res == xx
//}

/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	res := true
	arrs := make([]int, 0)
	for head != nil {
		arrs = append(arrs, head.Val)
		head = head.Next
	}
	if len(arrs) < 1 {
		return res
	}

	l, r := 0, len(arrs)-1
	for l < r {
		if l >= r {
			break
		}

		if arrs[l] != arrs[r] {
			res = false
			break
		}
		l++
		r--
	}
	return res
}
