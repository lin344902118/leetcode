package main

/*
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fake := &ListNode{}
	fake.Next = head
	left, right := fake, fake
	// right先走n步
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return head
}

func main() {

}
