package main

/*
206. 反转链表
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
	for right != nil {
		tmp := right.Next
		right.Next = left
		left = right
		right = tmp
	}
	return left
	// // 用一个链表存储每个节点的值
	// tmp := make([]int, 0)
	// node := head
	// for node != nil {
	// 	tmp = append(tmp, node.Val)
	// 	node = node.Next
	// }
	// // 重新遍历每个节点，然后倒过来赋值
	// node = head
	// for i := len(tmp) - 1; i >= 0; i-- {
	// 	node.Val = tmp[i]
	// 	node = node.Next
	// }
	// return head
}

func main() {

}
