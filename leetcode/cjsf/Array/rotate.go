package main

import "fmt"

/*
  旋转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	if k >= len(nums) {
		k = k % len(nums)
	}
	// 创建临时数组将后面的数存起来，然后先换后面的数，再换前面的数
	tmp := make([]int, 0, k)
	tmp = append(tmp, nums[len(nums)-k:]...)
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}
	return
}

/*
	if len(nums) < 2 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	for i := k; i < len(nums); i++ {
		nums[i] = tmp[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmp[len(nums)-k+i]
	}
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3 // [5,6,7,1,2,3,4]
	rotate(nums, k)
	fmt.Println(nums)
	nums = []int{-1, -100, 3, 99}
	k = 2 // [3,99,-1,-100]
	rotate(nums, k)
	fmt.Println(nums)
}
