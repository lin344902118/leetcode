package main

/*
104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftChildDepth := maxDepth(root.Left)
	rightChildDepth := maxDepth(root.Right)
	if leftChildDepth > rightChildDepth {
		return leftChildDepth + 1
	} else {
		return rightChildDepth + 1
	}
}

func main() {

}
