package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	curr, carry := head, 0
	x, y := 0, 0
	for l1 != nil || l2 != nil {
		x, y = 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		total := x + y + carry
		curr.Next = &ListNode{Val: total % 10}
		curr = curr.Next
		carry = total / 10
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head.Next
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	for head != nil {
		node := head.Next
		head.Next = pre
		pre = head
		head = node
	}
	return pre
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: -1,
	}
	curr, carry := head, 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		total := x + y + carry
		curr.Next = &ListNode{Val: total % 10}
		carry = total / 10
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head
}

type CardGroup struct {
	GroupType         int     `json:"groupType"`
	GroupName         string  `json:"groupName"`
	GroupURL          string  `json:"groupUrl"`
	GroupOrder        int     `json:"groupOrder"`
	BaseCount         int     `json:"baseCount"`
	NoLevelUp         bool    `json:"noLevelUp"`
	HaveRealNewCards  bool    `json:"haveRealNewCards"`
	CardsDetailURL    string  `json:"cardsDetailUrl"`
	TotalNum          int     `json:"totalNum"`
	GainNum           int     `json:"gainNum"`
	HaveNum           int     `json:"haveNum"`
	HaveShownThumbsUp bool    `json:"haveShownThumbsUp"`
	Cards             []*Card `json:"cards"`
}
type Card struct {
	CardStatus      int    `json:"cardStatus"`
	Toast           string `json:"toast"`
	CardId          int    `json:"cardId"`
	CardName        string `json:"cardName"`
	CardQuestion    string `json:"cardQuestion"`
	CardAnswer      string `json:"cardAnswer"`
	CardLevel       int    `json:"cardLevel"`
	CardLevelDesc   string `json:"cardLevelDesc"`
	CardNum         int    `json:"cardNum"`
	CardURL         string `json:"cardUrl"`
	CardGreyURL     string `json:"cardGreyUrl"`
	PersonURL       string `json:"personUrl"`
	PersonGreyURL   string `json:"personGreyUrl"`
	PreCardNum      int    `json:"preCardNum"`
	PreCardStar     int    `json:"preCardStar"`
	IsNewCard       bool   `json:"isNewCard"`
	IsRealNewCard   bool   `json:"isRealNewCard"`
	CardStar        int    `json:"cardStar"`
	CardNextRankNum int    `json:"cardNextRankNum"`
	FightingPoint   int    `json:"fightingPoint"`
	GainTime        int    `json:"gainTime"`
}

func main() {
	//strs := strings.Split(str ,",")
	//strs := strings.SplitN(str , "," , 0)
	//log.Println(len(strs))
	//log.Println(strs)
	//log.Println(strs[0])
	s1 := make([]int, 4, 10)
	s2 := s1[:50]
	log.Println(s2)
}
