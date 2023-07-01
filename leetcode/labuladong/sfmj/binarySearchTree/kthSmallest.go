package main

import "fmt"

/*
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	find := false
	res := -1
	var traverse func(root *TreeNode)
	// 中序遍历，寻找第k个
	traverse = func(root *TreeNode) {
		if root == nil || find {
			return
		}
		traverse(root.Left)
		k--
		if k == 0 {
			find = true
			res = root.Val
			return
		}
		traverse(root.Right)
		return
	}
	traverse(root)
	return res
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	fmt.Println(kthSmallest(root, 1))
}
