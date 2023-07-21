package main

import (
	"fmt"
	"math"
)

/*
124. 二叉树中的最大路径和
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 根节点大于最大值，更新最大值
		if root.Val > max {
			max = root.Val
		}
		var leftSum, rightSum int
		if root.Left != nil {
			// 左子树最大值
			leftSum = traverse(root.Left)
			// 左子树最大值大于最大值，更新最大值
			if leftSum > max {
				max = leftSum
			}
		}
		if root.Right != nil {
			// 右子树最大值
			rightSum = traverse(root.Right)
			// 右子树最大值大于最大值，更新最大值
			if rightSum > max {
				max = rightSum
			}
		}
		// 更新最大值
		if leftSum+root.Val > max {
			max = leftSum + root.Val
		}
		if rightSum+root.Val > max {
			max = rightSum + root.Val
		}
		if leftSum+root.Val+rightSum > max {
			max = leftSum + root.Val + rightSum
		}
		// 左右子树最大值都小于0，返回根节点自己
		if leftSum <= 0 && rightSum <= 0 {
			return root.Val
		} else {
			// 返回左子树和右子树最大值与根节点的和
			if leftSum > rightSum {
				return leftSum + root.Val
			} else {
				return rightSum + root.Val
			}
		}
	}
	traverse(root)
	return max
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(maxPathSum(root)) // 6
	root = &TreeNode{
		Val: -10,
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
	fmt.Println(maxPathSum(root)) // 42
	root = &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: -1,
			},
		},
	}
	fmt.Println(maxPathSum(root)) // 35
	root = &TreeNode{
		Val: -3,
	}
	fmt.Println(maxPathSum(root)) // -3
	root = &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: -6,
			},
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: -3,
			},
			Right: &TreeNode{
				Val: -6,
			},
		},
	}
	fmt.Println(maxPathSum(root)) // 10
	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(maxPathSum(root)) // 48
	root = &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: -6,
						Left: &TreeNode{
							Val: -6,
						},
					},
					Right: &TreeNode{
						Val: -6,
					},
				},
			},
		},
	}
	fmt.Println(maxPathSum(root)) // 16
}
