package main

import "fmt"

/*
96. 不同的二叉搜索树
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？
返回满足题意的二叉搜索树的种数。
*/

func numTrees(n int) int {
	tmp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		tmp[i] = make([]int, n+1)
	}
	var count func(low, high int) int
	count = func(low, high int) int {
		if low > high {
			return 1
		}
		if tmp[low][high] != 0 {
			return tmp[low][high]
		}
		res := 0
		// 总数等于从左到右固定每个数得到的左右数乘积的总和
		for mid := low; mid <= high; mid++ {
			left := count(low, mid-1)
			right := count(mid+1, high)
			res += left * right
		}
		tmp[low][high] = res
		return res
	}
	return count(1, n)
}

func main() {
	n := 3
	fmt.Println(numTrees(n))
}
