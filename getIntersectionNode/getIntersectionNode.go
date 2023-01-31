package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}


func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	curA , curB := headA , headB
	for curA != curB{
		if curA == nil{
			curA = headB
		}else{
			curA = curA.Next
		}

		if curB == nil{
			curB = headA
		}else{
			curB = curB.Next
		}
	}
	return curA
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA , curB := headA , headB
	for curA != curB{
		if curA == nil{
			curA = headB
		}else {
			curA = curA.Next
		}

		if curB == nil{
			curB = headA
		}else{
			curB = curB.Next
		}
	}
	return curA
}