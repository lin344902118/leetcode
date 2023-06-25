package main

import "fmt"

/*
327. 区间和的个数
给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper]
（包含 lower 和 upper）之内的 区间和的个数 。
区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
*/

// todo 思路未理清
func countRangeSum(nums []int, lower int, upper int) int {
	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		n := len(arr)
		if n <= 1 {
			return 0
		}

		n1 := append([]int(nil), arr[:n/2]...)
		n2 := append([]int(nil), arr[n/2:]...)
		cnt := mergeCount(n1) + mergeCount(n2) // 递归完毕后，n1 和 n2 均为有序

		// 统计下标对的数量
		l, r := 0, 0
		for _, v := range n1 {
			for l < len(n2) && n2[l]-v < lower {
				l++
			}
			for r < len(n2) && n2[r]-v <= upper {
				r++
			}
			cnt += r - l
		}

		// n1 和 n2 归并填入 arr
		p1, p2 := 0, 0
		for i := range arr {
			if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
				arr[i] = n1[p1]
				p1++
			} else {
				arr[i] = n2[p2]
				p2++
			}
		}
		return cnt
	}

	prefixSum := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSum[i+1] = prefixSum[i] + v
	}
	return mergeCount(prefixSum)
}

func main() {
	nums := []int{-2, 5, -1}
	fmt.Println(countRangeSum(nums, -2, 2))
}
