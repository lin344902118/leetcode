package main

import "fmt"

/*
300. 最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/

// func lengthOfLIS(nums []int) int {
// 	// 规定dp是指i位置的最长递增子序列，即前面有几个数小于当前数+1的个数
// 	dp := make([]int, len(nums))
// 	dp[0] = 1
// 	max := 1
// 	for i := 1; i < len(nums); i++ {
// 		//找到小于当前数的最大的dp
// 		dp[i] = 1
// 		for j := 0; j < i; j++ {
// 			if nums[j] < nums[i] {
// 				if dp[i] < dp[j]+1 {
// 					dp[i] = dp[j] + 1
// 				}
// 			}
// 		}
// 		if dp[i] > max {
// 			max = dp[i]
// 		}
// 	}
// 	return max
// }

func lengthOfLIS(nums []int) int {
	top := make([]int, len(nums))
	piles := 0
	for i := 0; i < len(nums); i++ {
		poker := nums[i]
		left := 0
		right := piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums = []int{0, 1, 0, 3, 2, 3}
	nums = []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}
