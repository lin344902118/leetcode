package main

import (
	"fmt"
)

/*
105. 从前序与中序遍历序列构造二叉树
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
 inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 根节点肯定是先序遍历的第一个节点。
	rootIndex := 0
	// 根据根节点将中序遍历分为两部分，在根节点之前的是左子树，之后的是右子树
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	// 左子树的节点数量
	leftNum := len(inorder[:rootIndex])
	left := buildTree(preorder[1:leftNum+1], inorder[:rootIndex])
	right := buildTree(preorder[leftNum+1:], inorder[rootIndex+1:])
	return &TreeNode{
		Val:   preorder[0],
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
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	preorderPrintTree(root)
}
