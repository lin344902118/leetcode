package main

import "fmt"

/*
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fake := &ListNode{
		Val:  -1,
		Next: head,
	}
	left := fake
	// 题目确保left不会为nil
	for i := 0; i < n; i++ {
		left = left.Next
	}
	right := fake
	for left.Next != nil {
		left = left.Next
		right = right.Next
	}
	// 此时left即为倒数第N-1个节点
	right.Next = right.Next.Next
	return fake.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	res := removeNthFromEnd(head, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
