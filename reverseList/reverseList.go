package reverseList

// lc 反转链表
type ListNode struct {
	Val  int
	Next *ListNode
}



func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList(head *ListNode) *ListNode{
	if head == nil{
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/*

sql题：给定一张消费流水表 flow，字段如下
id   name  product        unit_price       num    time
1   小明       A               5             3    2021-08-20 19:00:00
2   小宇       B               2             7    2021-08-20 19:00:00
3   小明       C               1            10   2021-08-20 19:00:00
…   …         …              …           …  ……
1、 统计在今年3月份期间，消费金额超过2000的同学
2、 统计在今年3月份期间，消费排名前3名的用户。
 */

/*
给定1,2,3,4,5,6,7,8,9固定9个数据，让任意5个数字组合成不重复的5位数，组合的数字也不能重复，
比如12345、23456、56789，将所有组合的数字输出，同时统计组合数字的数量并输出。
 */
