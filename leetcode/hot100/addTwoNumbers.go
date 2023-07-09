package main

/*
2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，
并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fake := &ListNode{
		Val: -1,
	}
	tmp := fake
	add := 0
	for l1 != nil || l2 != nil || add > 0 {
		val := add
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		add = val / 10
		val = val % 10
		tmp.Next = &ListNode{
			Val: val,
		}
		tmp = tmp.Next
	}
	return fake.Next
}

func main() {
}
