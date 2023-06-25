package main

import "fmt"

/*
654. 最大二叉树
给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:

创建一个根节点，其值为 nums 中的最大值。
递归地在最大值 左边 的 子数组前缀上 构建左子树。
递归地在最大值 右边 的 子数组后缀上 构建右子树。
返回 nums 构建的 最大二叉树 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[flag] {
			flag = i
		}
	}
	left := constructMaximumBinaryTree(nums[:flag])
	right := constructMaximumBinaryTree(nums[flag+1:])
	root := &TreeNode{
		Val:   nums[flag],
		Left:  left,
		Right: right,
	}
	return root
}

func preOrderPrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrderPrintTree(root.Left)
	preOrderPrintTree(root.Right)
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	preOrderPrintTree(root)
}
