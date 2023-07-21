package main

import "fmt"

/*
136. 只出现一次的数字
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，
其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/

func singleNumber(nums []int) int {
	// // map
	// m := make(map[int]int, 0)
	// for i := 0; i < len(nums); i++ {
	// 	_, e := m[nums[i]]
	// 	if e {
	// 		m[nums[i]]++
	// 	} else {
	// 		m[nums[i]] = 1
	// 	}
	// }
	// for num, count := range m {
	// 	if count == 1 {
	// 		return num
	// 	}
	// }
	// return -1
	if len(nums) == 0 {
		return 0
	}
	// 异或
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}
