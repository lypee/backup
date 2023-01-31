package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	reorderList(head)
	for head.Next != nil {
		log.Println(head.Val)
		head = head.Next
	}
}
func reorderList(head *ListNode) {
	copyHead := head
	list := make([]int, 0)
	length := 0
	for copyHead != nil {
		list = append(list, copyHead.Val)
		length++
		copyHead = copyHead.Next
	}
	newList := &ListNode{
		Val:  0,
		Next: nil,
	}
	bk := newList
	var val int
	left, right := 0, len(list)-1
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			val = list[left]
			left++
		} else {
			val = list[right]
			right--
		}
		//log.Printf("val :[%+v]", val)
		newList.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		newList = newList.Next
	}
	//log.Printf("val :[%+v]", val)
	head = bk.Next
}

func findMidNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func copyList(head *ListNode) *ListNode {
	cur := head
	back := cur
	for head != nil {
		cur.Next = &ListNode{
			Val: head.Val,
		}
		cur = cur.Next
		head = head.Next
	}
	return back.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func rev(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
