package main

import (
	"fmt"
	"math"
)

/*
1373. 二叉搜索子树的最大键值和
给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。

二叉搜索树的定义如下：

任意节点的左子树中的键值都 小于 此节点的键值。
任意节点的右子树中的键值都 大于 此节点的键值。
任意节点的左子树和右子树都是二叉搜索树。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	// 需要判断左右子树是否为BST，左子树最大值，右子树最小值和根节点的大小，以及左右子树总和
	ans := 0
	// 用一个长度为4的数组记录下需要的信息
	// 0表示是否是BST，1表示最小值，2表示最大值。3表示总和
	var dfs func(root *TreeNode) [4]int
	dfs = func(root *TreeNode) [4]int {
		if root == nil {
			return [4]int{1, math.MaxInt32, math.MinInt32, 0}
		}
		l, r := dfs(root.Left), dfs(root.Right)
		if l[0] == 1 && r[0] == 1 && l[2] < root.Val && r[1] > root.Val {
			s := l[3] + r[3] + root.Val
			ans = max(ans, s)
			return [4]int{1, min(l[1], root.Val), max(r[2], root.Val), s}
		}
		return [4]int{}
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Println(maxSumBST(root))
	root = &TreeNode{
		Val: -4,
		Left: &TreeNode{
			Val: -2,
		},
		Right: &TreeNode{
			Val: -5,
		},
	}
	fmt.Println(maxSumBST(root))
	root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(maxSumBST(root))
	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
	}
	fmt.Println(maxSumBST(root))
}
