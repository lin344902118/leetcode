package main

import "fmt"

/*
445. 两数相加 II
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	queue1 := make([]int, 0)
	queue2 := make([]int, 0)
	for l1 != nil {
		queue1 = append(queue1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		queue2 = append(queue2, l2.Val)
		l2 = l2.Next
	}
	fake := &ListNode{
		Val: -1,
	}
	add := 0
	for len(queue1) != 0 || len(queue2) != 0 || add > 0 {
		val := add
		if len(queue1) != 0 {
			val += queue1[len(queue1)-1]
			queue1 = queue1[:len(queue1)-1]
		}
		if len(queue2) != 0 {
			val += queue2[len(queue2)-1]
			queue2 = queue2[:len(queue2)-1]
		}
		add = val / 10
		val = val % 10
		node := &ListNode{
			Val: val,
		}
		node.Next = fake.Next
		fake.Next = node
	}
	return fake.Next
}

func printTree(root *ListNode) {
	tmp := root
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

func main() {
	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Println(addTwoNumbers(l1, l2))
}
