package main

import "fmt"

/*
114. 二叉树展开为链表
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	// if root == nil {
	// 	return
	// }
	// // 获取先序遍历结果，然后修改
	// res := make([]int, 0)
	// var preOrderTraverse func(root *TreeNode)
	// preOrderTraverse = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	res = append(res, root.Val)
	// 	preOrderTraverse(root.Left)
	// 	preOrderTraverse(root.Right)
	// }
	// preOrderTraverse(root)
	// tmp := root
	// for i := 1; i < len(res); i++ {
	// 	tmp.Left = nil
	// 	tmp.Right = &TreeNode{
	// 		Val: res[i],
	// 	}
	// 	tmp = tmp.Right
	// }
	// return
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left
	tmp := root
	for tmp.Right != nil {
		tmp = tmp.Right
	}
	tmp.Right = right
	return
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
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(root)
	preOrderTraverse(root)
}
