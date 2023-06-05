package main

import (
	"sort"
)

/*
  合并两个有序数组
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，
nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，
应忽略。nums2 的长度为 n 。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 插入后排序
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
	// 双指针做法
	// if n == 0 {
	// 	return
	// }
	// tmp := make([]int, m)
	// copy(tmp, nums1[:m])
	// left, right := 0, 0
	// for left < m && right < n {
	// 	if tmp[left] <= nums2[right] {
	// 		nums1[left+right] = tmp[left]
	// 		left++
	// 	} else {
	// 		nums1[left+right] = nums2[right]
	// 		right++
	// 	}
	// }
	// if left == m {
	// 	for i := right; i < n; i++ {
	// 		nums1[m+i] = nums2[i]
	// 	}
	// }
	// if right == n {
	// 	for i := left; i < m; i++ {
	// 		nums1[n+i] = tmp[i]
	// 	}
	// }
	// return
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	nums1 = []int{1, 2, 4, 5, 6, 0}
	m = 5
	nums2 = []int{3}
	n = 1
	merge(nums1, m, nums2, n)
}
