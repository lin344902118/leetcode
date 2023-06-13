package main

/*
86. 分隔链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，
使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small, big := &ListNode{}, &ListNode{}
	tmp1, tmp2 := small, big
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}
		// 断开之前的连接
		tmp := head.Next
		head.Next = nil
		head = tmp
	}
	small.Next = tmp2.Next
	return tmp1.Next
}

func main() {

}
