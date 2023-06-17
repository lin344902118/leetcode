package main

import "fmt"

/*
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。
*/

func maxSlidingWindow(nums []int, k int) []int {
	// 单调队列
	queue := make([]int, 0)
	res := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		// 将前k-1个加入
		if i < k-1 {
			// 小于当前入队元素的弹出
			for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
				queue = queue[:len(queue)-1]
			}
			// 将元素加入
			queue = append(queue, nums[i])
		} else {
			// 小于当前入队元素的弹出
			for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
				queue = queue[:len(queue)-1]
			}
			// 将元素加入
			queue = append(queue, nums[i])
			// 队头即最大值
			res = append(res, queue[0])
			// 如果队头元素时窗口第一个值，需要删除。如果不是则不处理
			if queue[0] == nums[i-k+1] {
				queue = queue[1:]
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
	nums = []int{1}
	k = 1
	fmt.Println(maxSlidingWindow(nums, k))
}
