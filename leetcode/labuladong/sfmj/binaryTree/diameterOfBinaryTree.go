package main

import "fmt"

/*
543. 二叉树的直径
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := getmaxDepth(root, 0)
	return res
}

func getmaxDepth(root *TreeNode, res int) (int, int) {
	if root == nil {
		return 0, res
	}
	leftMaxDepth, res := getmaxDepth(root.Left, res)
	rightMaxDepth, res := getmaxDepth(root.Right, res)
	if leftMaxDepth+rightMaxDepth > res {
		res = leftMaxDepth + rightMaxDepth
	}
	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth + 1, res
	} else {
		return rightMaxDepth + 1, res
	}
}

func main() {
	root := &TreeNode{
		Left: &TreeNode{
			Val: 2,
		},
		Val: 1,
	}
	fmt.Println(diameterOfBinaryTree(root))
}
