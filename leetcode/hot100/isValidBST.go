package main

import "fmt"

/*
98. 验证二叉搜索树
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

func isValidBST(root *TreeNode) bool {
	// // 获取中序遍历结果
	// res := make([]int, 0)
	// var traverse func(root *TreeNode)
	// traverse = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	traverse(root.Left)
	// 	res = append(res, root.Val)
	// 	traverse(root.Right)
	// }
	// traverse(root)
	// for i := 0; i < len(res)-1; i++ {
	// 	if res[i] >= res[i+1] {
	// 		return false
	// 	}
	// }
	// return true
	var isBST func(root, min, max *TreeNode) bool
	isBST = func(root, min, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		return isBST(root.Left, min, root) && isBST(root.Right, root, max)
	}
	return isBST(root, nil, nil)
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(isValidBST(root))
}
