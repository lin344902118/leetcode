package main

import "fmt"

/*
503. 下一个更大元素 II
给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），
返回 nums 中每个元素的 下一个更大元素 。
数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，
这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。
*/

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	// 遍历两次nums
	for i := 2*len(nums) - 1; i >= 0; i-- {
		num := nums[i%len(nums)]
		// 将小于等于num的数弹出
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if i < len(nums) {
			if len(stack) == 0 {
				res[i] = -1
			} else {
				res[i] = stack[len(stack)-1]
			}
		}
		stack = append(stack, num)
	}
	return res
}

func main() {
	nums := []int{1, 2, 1}
	// fmt.Println(nextGreaterElements(nums))
	// nums = []int{1, 2, 3, 4, 3}
	// fmt.Println(nextGreaterElements(nums))
	nums = []int{100, 1, 11, 1, 120, 111, 123, 1, -1, -100}
	// 120 11 120 120 123 123 -1 100 100 100
	fmt.Println(nextGreaterElements(nums))
}
