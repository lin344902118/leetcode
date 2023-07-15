package main

import (
	"fmt"
	"sort"
)

/*
169. 多数元素
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

func majorityElement(nums []int) int {
	// 思路一
	// if len(nums) == 1 {
	// 	return nums[0]
	// }
	// cache := make(map[int]int, 0)
	// for i := 0; i < len(nums); i++ {
	// 	count, e := cache[nums[i]]
	// 	if e {
	// 		count++
	// 		if count > len(nums)/2 {
	// 			return nums[i]
	// 		}
	// 		cache[nums[i]] = count
	// 	} else {
	// 		cache[nums[i]] = 1
	// 	}
	// }
	// 思路二
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	nums := []int{3, 2, 3}
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
