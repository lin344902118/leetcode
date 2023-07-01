package main

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
	return traverse(root, nil, nil)
}

func traverse(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return traverse(root.Left, min, root) && traverse(root.Right, root, max)
}

// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	res := traverse(root)
// 	for i := 0; i < len(res)-1; i++ {
// 		if res[i] >= res[i+1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func traverse(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	if root == nil {
// 		return res
// 	}
// 	res = append(res, traverse(root.Left)...)
// 	res = append(res, root.Val)
// 	res = append(res, traverse(root.Right)...)
// 	return res
// }

func main() {

}
