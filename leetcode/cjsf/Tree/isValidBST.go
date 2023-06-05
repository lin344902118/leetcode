package main

import (
	"fmt"
	"math"
)

/*
  验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	res := make([]int, 0)
// 	res = middleBST(root, res)
// 	for i := 0; i < len(res)-1; i++ {
// 		if res[i] >= res[i+1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func middleBST(root *TreeNode, res []int) []int {
// 	if root == nil {
// 		return []int{}
// 	}
// 	res = append(res, middleBST(root.Left, []int{})...)
// 	res = append(res, root.Val)
// 	res = append(res, middleBST(root.Right, []int{})...)
// 	return res
// }

func isValidBST(root *TreeNode) bool {
	return validBST(root, math.MaxInt64, math.MinInt64)
}

func validBST(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return validBST(root.Left, root.Val, min) && validBST(root.Right, max, root.Val)
}

func main() {
	head := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(isValidBST(head)) // true
	head = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(isValidBST(head)) // false
	head = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	fmt.Println(isValidBST(head)) // false

}
