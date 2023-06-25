package main

import (
	"fmt"
	"strconv"
)

/*
652. 寻找重复的子树
给你一棵二叉树的根节点 root ，返回所有 重复的子树 。

对于同一类的重复子树，你只需要返回其中任意 一棵 的根结点即可。

如果两棵树具有 相同的结构 和 相同的结点值 ，则认为二者是 重复 的。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	counts := make(map[string]int, 0)
	res := make([]*TreeNode, 0)
	var traverse func(*TreeNode) string
	traverse = func(tn *TreeNode) string {
		if tn == nil {
			return "#"
		}
		left := traverse(tn.Left)
		right := traverse(tn.Right)
		subTree := left + "," + right + "," + strconv.Itoa(tn.Val)
		count, exist := counts[subTree]
		if exist && count == 1 {
			res = append(res, tn)
		}
		counts[subTree]++
		return subTree
	}
	traverse(root)
	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 11,
			},
		},
		Right: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}
	res := findDuplicateSubtrees(root)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].Val)
	}
}
