package detectCycle

type ListNode struct {
	Val  int
	Next *ListNode
}


func detectCycle(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	slow ,fast := head , head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}


func detectCycle2(head *ListNode) *ListNode{
	if head == nil{
		return nil
	}
	slow , fast := head , head
	for fast != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			for slow != head{
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}