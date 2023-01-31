package main

import (
	"context"
	"log"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0,
		Next: head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		second = second.Next
		first = first.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	if second.Next != nil {
		second.Next = second.Next.Next
	}
	return dummy.Next
}

func main() {
	pCtx := context.Background()
	ctrlCtx, _ := context.WithTimeout(pCtx, time.Duration(5)*time.Second)
	ch := make(chan struct{}, 1)
	go writeToChan(ch)

	for {
		select {
		case <-ctrlCtx.Done():
			log.Println("timeout!")
			return
		case data, ok := <-ch:
			if ok {
				log.Printf("read value:[%+v]", data)
				break
			} else {
				log.Printf("not read value:[%+v]", data)
			}

		}
	}
}

func writeToChan(ch chan struct{}) {
	time.Sleep(6 * time.Second)
	ch <- struct{}{}
}
