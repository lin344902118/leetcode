package main

import "fmt"

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
	// res := 0
	// if root == nil {
	// 	return res
	// }
	// // BFS
	// queue := make([]*TreeNode, 0)
	// queue = append(queue, root)
	// for len(queue) != 0 {
	// 	num := len(queue)
	// 	for i := 0; i < num; i++ {
	// 		first := queue[0]
	// 		queue = queue[1:]
	// 		if first.Left != nil {
	// 			queue = append(queue, first.Left)
	// 		}
	// 		if first.Right != nil {
	// 			queue = append(queue, first.Right)
	// 		}
	// 	}
	// 	res++
	// }
	// return res
	// // DFS
	// level := 1
	// var traverse func(root *TreeNode, level int) int
	// traverse = func(root *TreeNode, level int) int {
	// 	if root == nil {
	// 		return level - 1
	// 	}
	// 	leftLevel := traverse(root.Left, level+1)
	// 	rightLevel := traverse(root.Right, level+1)
	// 	if leftLevel > rightLevel {
	// 		return leftLevel
	// 	} else {
	// 		return rightLevel
	// 	}
	// }
	// return traverse(root, level)

	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(maxDepth(root))
}
