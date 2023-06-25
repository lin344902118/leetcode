package main

/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)
	return root
}

// 将二叉树当成三叉树遍历
func traverse(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	// 顶层直接连接
	left.Next = right
	// 连接左子树左右子树
	traverse(left.Left, left.Right)
	// 连接右子树左右子树
	traverse(right.Left, right.Right)
	// 连接左子树右子树和右子树左子树
	traverse(left.Right, right.Left)
}

func main() {

}
