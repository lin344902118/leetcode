package main

import "fmt"

/*
106. 从中序与后序遍历序列构造二叉树
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历，
postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 后序遍历的最后一个值是根节点，根据这个值将中序遍历分离
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			rootIndex = i
			break
		}
	}
	leftNum := len(inorder[:rootIndex])
	left := buildTree(inorder[:rootIndex], postorder[:leftNum])
	right := buildTree(inorder[rootIndex+1:], postorder[leftNum:len(postorder)-1])
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  left,
		Right: right,
	}
}

func preorderPrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderPrintTree(root.Left)
	preorderPrintTree(root.Right)
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	preorderPrintTree(root)
}
