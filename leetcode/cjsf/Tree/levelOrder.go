package main

import "fmt"

/*
  二叉树的层序遍历
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 创建队列用来保存每层的节点
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		// 弹出每层的节点，并加上每层的子节点
		levelNum := len(queue)
		subRes := make([]int, 0)
		for i := 0; i < levelNum; i++ {
			node := queue[0]
			subRes = append(subRes, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, subRes)
	}
	return res
}

// DFS
// func levelOrder(root *TreeNode) [][]int {
// 	res := make([][]int, 0)
// 	helper(&res, root, 0)
// 	return res
// }

// func helper(res *[][]int, root *TreeNode, level int) {
// 	if root == nil {
// 		return
// 	}
// 	if level >= len(*res) {
// 		*res = append(*res, []int{})
// 	}
// 	(*res)[level] = append((*res)[level], root.Val)
// 	helper(res, root.Left, level+1)
// 	helper(res, root.Right, level+1)
// 	return
// }

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 2,
			},
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
