package main

/*
  反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left := head
	right := head.Next
	left.Next = nil
	for right.Next != nil {
		tmp := right.Next
		right.Next = left
		left = right
		right = tmp
	}
	right.Next = left
	return right
}

func main() {

}
