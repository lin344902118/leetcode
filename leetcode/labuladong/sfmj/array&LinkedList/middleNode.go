package main

/*
876. 链表的中间结点
给你单链表的头结点 head ，请你找出并返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}
	return slow
}

func main() {

}
