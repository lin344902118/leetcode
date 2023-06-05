package main

import "fmt"

/*
  对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isSymmetric(root)) // true
	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Println(isSymmetric(root)) // false
	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
	}
	fmt.Println(isSymmetric(root)) // false
}
