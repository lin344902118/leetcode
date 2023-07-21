package main

import "fmt"

/*
101. 对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// // 中序遍历，遍历的结果需要加上深度，否则不唯一
	// res := make([]string, 0)
	// var midTraverse func(root *TreeNode, depth int)
	// midTraverse = func(root *TreeNode, depth int) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	midTraverse(root.Left, depth+1)
	// 	res = append(res, fmt.Sprintf("%d_%d", root.Val, depth))
	// 	midTraverse(root.Right, depth+1)
	// }
	// midTraverse(root, 1)
	// left := 0
	// right := len(res) - 1
	// for left <= right {
	// 	if res[left] != res[right] {
	// 		return false
	// 	}
	// 	left++
	// 	right--
	// }
	// return true
	// 递归判断
	if root == nil {
		return true
	}
	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root.Left, root.Right)
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
	fmt.Println(isSymmetric(root))
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
	fmt.Println(isSymmetric(root))
}
