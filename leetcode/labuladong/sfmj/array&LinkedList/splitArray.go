package main

import "fmt"

/*
410. 分割数组的最大值
给定一个非负整数数组 nums 和一个整数 m ，你需要将这个数组分成 m 个非空的连续子数组。
设计一个算法使得这 m 个子数组各自和的最大值最小。
*/

func splitArray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	minNum := 0
	maxNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > minNum {
			minNum = nums[i]
		}
		maxNum += nums[i]
	}
	for minNum < maxNum {
		mid := minNum + ((maxNum - minNum) >> 1)
		n := getMax(nums, mid)
		if n == k {
			maxNum = mid
		} else if n < k {
			maxNum = mid
		} else if n > k {
			minNum = mid + 1
		}
	}
	return minNum
}

// 给一个数组和最大值，按最大值能分成几个数组
func getMax(nums []int, max int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	m := max
	for i := 0; i < len(nums); i++ {
		if m < nums[i] {
			count++
			m = max
		}
		m -= nums[i]
	}
	return count + 1
}

func main() {
	nums := []int{7, 2, 5, 10, 8}
	m := 2
	fmt.Println(splitArray(nums, m))
	nums = []int{1, 2, 3, 4, 5}
	m = 2
	fmt.Println(splitArray(nums, m))
	nums = []int{1, 4, 4}
	m = 3
	fmt.Println(splitArray(nums, m))
}
