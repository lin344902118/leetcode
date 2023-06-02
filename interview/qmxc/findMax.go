package main

import "fmt"

/*
写一个函数，要求输入一个Node根节点，搜索该节点所有路径，
路径下节点的value为偶数则加上该value，为奇数则减去该value，
计算并输出所有路径的最大值。
*/
type Node struct {
	value int
	left  *Node
	right *Node
}

func FindMax(node *Node) int {
	if node == nil {
		return 0
	}
	num := 0
	if node.value%2 == 0 {
		num += node.value
	} else {
		num -= node.value
	}
	left := FindMax(node.left)
	right := FindMax(node.right)
	if num+left > num+right {
		return num + left
	} else {
		return num + right
	}
}

func main() {
	left2 := &Node{
		value: 8,
		left:  nil,
		right: nil,
	}
	left1 := &Node{
		value: 1,
		left:  left2,
		right: nil,
	}
	right1 := &Node{
		value: 4,
		left:  nil,
		right: nil,
	}
	root := &Node{
		value: 2,
		left:  left1,
		right: right1,
	}
	fmt.Println(FindMax(root))
}
