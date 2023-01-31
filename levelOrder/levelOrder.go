package main

import (
	"encoding/json"
	"log"
)

func main() {
	s := []int{6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18}
	ss, _ := json.Marshal(&s)
	log.Println(ss)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	s := make([]*TreeNode, 0)
	s = append(s, root)
	for len(s) != 0 {
		n := len(s)
		sList := []int{}
		for i := 0; i < n; i++ {
			node := s[0]
			s = s[1:]
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
			sList = append(sList, node.Val)
		}
		res = append(res, sList)
	}
	return res
}
