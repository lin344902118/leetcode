package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
297. 二叉树的序列化与反序列化
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。
你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 先序遍历
	sb := strings.Builder{}
	if root == nil {
		sb.WriteByte('#')
		sb.WriteByte(',')
		return sb.String()
	}
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteByte(',')
	sb.WriteString(this.serialize(root.Left))
	sb.WriteString(this.serialize(root.Right))
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	datas := strings.Split(data, ",")
	var helper func() *TreeNode
	helper = func() *TreeNode {
		if len(datas) == 0 {
			return nil
		}
		if datas[0] == "#" {
			datas = datas[1:]
			return nil
		}
		rootVal, _ := strconv.Atoi(datas[0])
		datas = datas[1:]
		return &TreeNode{Val: rootVal, Left: helper(), Right: helper()}
	}
	return helper()
}

func preorderPrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderPrintTree(root.Left)
	preorderPrintTree(root.Right)
}

func main() {
	// 1,2,3,4,5
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	// root = &TreeNode{
	// 	Val: 1,
	// 	Right: &TreeNode{
	// 		Val: 2,
	// 	},
	// }
	c := Constructor()
	s := c.serialize(root)
	fmt.Println(s)
	node := c.deserialize(s)
	preorderPrintTree(node)
}
