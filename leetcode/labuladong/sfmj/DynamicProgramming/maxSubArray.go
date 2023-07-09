package main

import (
	"fmt"
)

/*
53. 最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
*/

func maxSubArray(nums []int) int {
	// 滑动窗口
	// left, right := 0, 0
	// windowSum := 0
	// maxSum := math.MinInt32
	// for right < len(nums) {
	// 	windowSum += nums[right]
	// 	right++
	// 	if windowSum > maxSum {
	// 		maxSum = windowSum
	// 	}
	// 	for windowSum < 0 {
	// 		windowSum -= nums[left]
	// 		left++
	// 	}
	// }
	// return maxSum
	dp_0 := nums[0]
	dp_1 := 0
	max := dp_0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]+dp_0 {
			dp_1 = nums[i]
		} else {
			dp_1 = dp_0 + nums[i]
		}
		dp_0 = dp_1
		if dp_1 > max {
			max = dp_1
		}
	}
	return max
}

func main() {
	nums := []int{-3, -1, -2}
	fmt.Println(maxSubArray(nums))
}
