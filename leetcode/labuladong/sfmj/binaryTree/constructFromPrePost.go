package main

import "fmt"

/*
889. 根据前序和后序遍历构造二叉树
给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵树的后序遍历，重构并返回二叉树。

如果存在多个答案，您可以返回其中 任何 一个。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}
	if len(preorder) == 1 || len(postorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}
	// 假设前序遍历的第⼆个元素是左⼦树的根节点，但实际上左⼦树有可能是空指针
	// 先序遍历的第0个元素是根节点，假设第1个是左子树。根据这个将后序遍历分离，然后构造子树
	leftIndex := 0
	for i := 0; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			leftIndex = i
			break
		}
	}
	leftNum := len(postorder[:leftIndex+1])
	left := constructFromPrePost(preorder[1:1+leftNum], postorder[:leftIndex+1])
	right := constructFromPrePost(preorder[1+leftNum:], postorder[leftIndex+1:len(postorder)-1])
	return &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
}

func inorderPrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	inorderPrintTree(root.Left)
	fmt.Println(root.Val)
	inorderPrintTree(root.Right)
}

func main() {
	preOrder := []int{1, 2, 4, 5, 3, 6, 7}
	postOrder := []int{4, 5, 2, 6, 7, 3, 1}
	root := constructFromPrePost(preOrder, postOrder)
	inorderPrintTree(root)
}
