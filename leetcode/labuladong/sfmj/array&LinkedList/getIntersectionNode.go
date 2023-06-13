package main

/*
160. 相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

图示两个链表在节点 c1 开始相交：

题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	left := headA
	right := headB
	// 将两个环相连，如果相遇则表示有交叉，没有相遇则没有交叉
	for left != right {
		if left == nil {
			left = headB
		} else {
			left = left.Next
		}
		if right == nil {
			right = headA
		} else {
			right = right.Next
		}
	}
	return right
}

func main() {

}
