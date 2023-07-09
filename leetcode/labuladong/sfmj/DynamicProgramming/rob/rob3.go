package main

/*
337. 打家劫舍 III
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，
聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	result := robTree(root)
	return max(result[0], result[1])
}

func robTree(root *TreeNode) []int {
	if root == nil {
		// 0表示抢劫，1表示不抢劫
		return []int{0, 0}
	}
	left := robTree(root.Left)
	right := robTree(root.Right)

	// 状态转移方程
	// dp[0] 表示不抢劫当前根节点，dp[1]表示抢劫当前节点
	// dp[0] = dp[left], dp[right]
	// dp[1] = root.Val + dp[left.left] + dp[left.right] + dp[right.left] + dp[right.right]
	// 当前节点加上不抢劫左右子树
	rb := root.Val + left[1] + right[1]
	// 抢劫左右子树
	urb := max(left[0], left[1]) + max(right[0], right[1])
	return []int{rb, urb}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
}
