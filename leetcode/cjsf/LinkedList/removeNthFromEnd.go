package main

/*
  删除链表的倒数第N个节点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fake := &ListNode{
		Val:  0,
		Next: head,
	}
	left := fake
	right := fake
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return fake.Next
}

func main() {

}
