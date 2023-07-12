package main

import "fmt"

/*
23. 合并 K 个升序链表
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type heap []int

// 交换两个节点的值
func (h heap) swap(i, j int) {
	if i < 0 || i > len(h) || j < 0 || j > len(h) {
		return
	}
	h[i], h[j] = h[j], h[i]
	return
}

// 比较两个节点的值
func (h heap) less(i, j int) bool {
	if i < 0 || i > len(h) || j < 0 || j > len(h) {
		return false
	}
	return h[i] < h[j]
}

// 节点向上移动 最小堆最后一个节点，需要跟父节点对比，找到一个父节点比他小的
func (h heap) up(i int) {
	if i < 0 || i > len(h) {
		return
	}
	for {
		// 获取父节点值
		f := (i - 1) / 2
		// 父节点小于等于该节点退出
		if i == f || h.less(f, i) {
			break
		}
		// 交换父子节点值
		h.swap(f, i)
		i = f
	}
	return
}

// 节点向下移动 最小堆节点，需要跟子节点对比，直到左节点小于该节点，右节点大于该节点
func (h heap) down(i int) {
	if i < 0 || i > len(h) {
		return
	}
	for {
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		// 当前节点为叶子节点
		if leftChild >= len(h) {
			break
		}
		// 获取最小子节点
		min := leftChild
		if rightChild < len(h) && h.less(rightChild, leftChild) {
			min = rightChild
		}
		// 当前节点已经小于最小子节点，直接退出
		if h.less(i, min) {
			break
		}
		h.swap(i, min)
		i = min
	}
	return
}

// 将节点添加进堆
func (h *heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

// 从堆中删除节点
func (h *heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}
	// 跟最后一个值交换
	n := len(*h) - 1
	h.swap(i, n)
	// 删除最后一个值
	x := (*h)[n]
	*h = (*h)[:n]
	// 上移或者下移
	if i == 0 || (*h)[i] >= (*h)[(i-1)/2] {
		// 父节点小于该节点，说明上面已经满足，下移
		h.down(i)
	} else {
		h.up(i)
	}
	return x, true
}

// 删除堆顶节点
func (h *heap) Pop() int {
	x, _ := h.Remove(0)
	return x
}

func NewHeap(arr []int) heap {
	h := heap(arr)
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
	return h
}

func mergeKLists(lists []*ListNode) *ListNode {
	arr := make([]int, 0)
	for i := 0; i < len(lists); i++ {
		for lists[i] != nil {
			arr = append(arr, lists[i].Val)
			lists[i] = lists[i].Next
		}
	}
	fake := &ListNode{}
	tmp := fake
	heap := NewHeap(arr)
	fmt.Println(heap)
	for len(heap) > 0 {
		node := &ListNode{
			Val: heap.Pop(),
		}
		tmp.Next = node
		tmp = tmp.Next
	}
	return fake.Next
}

func main() {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	node2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	node3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}
	lists := []*ListNode{node1, node2, node3}
	lists = []*ListNode{}
	root := mergeKLists(lists)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}
