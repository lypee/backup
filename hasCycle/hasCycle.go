package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}


func hasCycle2(head *ListNode) bool {
	if head == nil{
		return true
	}
	slow , fast := head , head.Next
	for fast != nil && fast.Next != nil{
		if slow == fast{
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
