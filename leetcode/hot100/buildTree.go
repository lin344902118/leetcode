package main

import "fmt"

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
	if len(preorder) != len(preorder) {
		return nil
	}
	// 先序遍历的第一个元素是根节点
	rootIndex := 0
	for i := rootIndex; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	// 1到rootIndex+1为左子树数量
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+rootIndex], inorder[:rootIndex]),
		Right: buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:]),
	}
	return root
}

func preOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrderTraverse(root.Left)
	preOrderTraverse(root.Right)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	preOrderTraverse(root)
}
