package main

import "fmt"

/*
102. 二叉树的层序遍历
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// // BFS
	// res := make([][]int, 0)
	// if root == nil {
	// 	return res
	// }
	// queue := make([]*TreeNode, 0)
	// queue = append(queue, root)
	// for len(queue) != 0 {
	// 	// 每一层的数量
	// 	num := len(queue)
	// 	tmp := make([]int, 0)
	// 	for i := 0; i < num; i++ {
	// 		first := queue[0]
	// 		tmp = append(tmp, first.Val)
	// 		queue = queue[1:]
	// 		if first.Left != nil {
	// 			queue = append(queue, first.Left)
	// 		}
	// 		if first.Right != nil {
	// 			queue = append(queue, first.Right)
	// 		}
	// 	}
	// 	res = append(res, tmp)
	// }
	// return res
	// DFS
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var helper func(root *TreeNode, level int)
	helper = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		helper(root.Left, level+1)
		helper(root.Right, level+1)
		return
	}
	helper(root, 0)
	return res
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
	fmt.Println(levelOrder(root))
}
