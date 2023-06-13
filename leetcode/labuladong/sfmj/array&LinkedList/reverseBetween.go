package main

import "fmt"

/*
92. 反转链表 II
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	tmp := make([]int, right-left+1)
	node := head
	for i := 1; i < left; i++ {
		node = node.Next
	}
	n := node
	for i := 0; i <= right-left; i++ {
		tmp[i] = n.Val
		n = n.Next
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		node.Val = tmp[i]
		node = node.Next
	}
	return head
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// // 虚拟头节点，避免边界问题
	// fake := &ListNode{
	// 	Next: head,
	// }
	// node := fake
	// for i := 1; i < left; i++ {
	// 	node = node.Next
	// }
	// pre := node
	// slow := pre.Next
	// fast := slow.Next
	// // 开始反转
	// for i := 0; i < right-left; i++ {
	// 	tmp := fast.Next
	// 	fast.Next = slow
	// 	slow = fast
	// 	fast = tmp
	// }
	// // 链接反转后的链表
	// pre.Next.Next = fast
	// pre.Next = slow
	// return fake.Next
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
	head = reverseBetween(head, 2, 4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
