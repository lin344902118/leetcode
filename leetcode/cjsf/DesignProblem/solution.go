package main

import (
	"math/rand"
)

/*
  打乱数组
给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。

实现 Solution class:

Solution(int[] nums) 使用整数数组 nums 初始化对象
int[] reset() 重设数组到它的初始状态并返回
int[] shuffle() 返回数组随机打乱后的结果
*/

type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	return Solution{
		data: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.data
}

func (this *Solution) Shuffle() []int {
	res := make([]int, len(this.data))
	copy(res, this.data)
	for i := 0; i < len(this.data); i++ {
		j := rand.Intn(i + 1)
		if i != j {
			res[i] ^= res[j]
			res[j] ^= res[i]
			res[i] ^= res[j]
		}
	}
	return res
}

func main() {
}
