package main

import "fmt"

/*
496. 下一个更大元素 I
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，
并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 单调栈加哈希表
	cache := make(map[int]int)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		// 将栈内小于等于的元素弹出
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		// 栈内没有元素，说明没有更大的，有的话取最后一个，用哈希表缓存
		if len(stack) == 0 {
			cache[num] = -1
		} else {
			cache[num] = stack[len(stack)-1]
		}
		stack = append(stack, num)
	}
	res := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res = append(res, cache[nums1[i]])
	}
	return res
	// 暴力
	// res := make([]int, 0, len(nums1))
	// for i := 0; i < len(nums1); i++ {
	// 	index := 0
	// 	for j := 0; j < len(nums2); j++ {
	// 		if nums2[j] == nums1[i] {
	// 			index = j
	// 			break
	// 		}
	// 	}
	// 	find := false
	// 	for k := index + 1; k < len(nums2); k++ {
	// 		if nums2[k] > nums2[index] {
	// 			find = true
	// 			res = append(res, nums2[k])
	// 			break
	// 		}
	// 	}
	// 	if !find {
	// 		res = append(res, -1)
	// 	}
	// }
	// return res
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
	nums1 = []int{2, 4}
	nums2 = []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
