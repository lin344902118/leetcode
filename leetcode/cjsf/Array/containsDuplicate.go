package main

import "fmt"

/*
  存在重复元素
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；
如果数组中每个元素互不相同，返回 false 。
*/

func containsDuplicate(nums []int) bool {
	cache := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		_, exist := cache[nums[i]]
		if exist {
			return true
		}
		cache[nums[i]] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1} // true
	fmt.Println(containsDuplicate(nums))
	nums = []int{1, 2, 3, 4} // false
	fmt.Println(containsDuplicate(nums))
	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2} // true
	fmt.Println(containsDuplicate(nums))
}
