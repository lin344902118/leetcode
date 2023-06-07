package main

import "fmt"

/*
  回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。
如果是，返回 true ；否则，返回 false 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 使用数组记录每个节点值，判断数组是否是链表
	// arr := make([]int, 0)
	// for head != nil {
	// 	arr = append(arr, head.Val)
	// 	head = head.Next
	// }
	// for i := 0; i < len(arr)/2; i++ {
	// 	if arr[i] != arr[len(arr)-1-i] {
	// 		return false
	// 	}
	// }
	// return true
	// 使用快慢指针找到链表中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}
	// 反转之后的链表
	slow = reverseList(slow)
	// 判断两个链表是否相等
	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}
	return true
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
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(isPalindrome(head))
}
