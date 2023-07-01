package main

/*
450. 删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，
并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
一般来说，删除节点可分为两个步骤：
首先找到需要删除的节点；
如果找到了，删除它。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		// 有四种情况，叶子节点左右子树都为空，左子树为空右子树不为空，右子树为空左子树不为空，左右子树都不为空
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 找到右子树最小值或者左子树最大值。
		minNode := getMin(root.Right)
		root.Right = deleteNode(root.Right, minNode.Val)
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
		// maxNode := getMax(root.Left)
		// root.Left = deleteNode(root.Left, maxNode.Val)
		// maxNode.Left = root.Left
		// maxNode.Right = root.Right
		// root = maxNode
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func getMax(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func main() {

}
