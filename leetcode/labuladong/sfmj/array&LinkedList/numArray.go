package main

import "fmt"

/*
303. 区域和检索 - 数组不可变
给定一个整数数组  nums，处理以下类型的多个查询:

计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，
包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
*/

type NumArray struct {
	nums   []int
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, 0, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		preSum = append(preSum, sum)
	}
	fmt.Println(preSum)
	return NumArray{
		nums:   []int{},
		preSum: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.preSum[right]
	} else {
		return this.preSum[right] - this.preSum[left-1]
	}
}

func main() {
	arr := []int{-2, 0, 3, -5, 2, -1}
	numArr := Constructor(arr)
	fmt.Println(numArr.SumRange(0, 2))
	fmt.Println(numArr.SumRange(2, 5))
	fmt.Println(numArr.SumRange(0, 5))
}
