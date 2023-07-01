package main

/*
95. 不同的二叉搜索树 II
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。
可以按 任意顺序 返回答案。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var help func(start, end int) []*TreeNode
	help = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		res := make([]*TreeNode, 0)
		// 从0-n作为根节点
		for i := start; i <= end; i++ {
			// 根节点左右子树
			leftTrees := help(start, i-1)
			rightTrees := help(i+1, end)
			for j := 0; j < len(leftTrees); j++ {
				for k := 0; k < len(rightTrees); k++ {
					root := &TreeNode{
						Val: i,
					}
					root.Left = leftTrees[j]
					root.Right = rightTrees[k]
					res = append(res, root)
				}
			}
		}
		return res
	}
	return help(1, n)
}

func main() {

}
