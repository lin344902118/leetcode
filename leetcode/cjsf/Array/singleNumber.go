package main

import "fmt"

/*
  只出现一次的数字
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {
	nums := []int{2, 2, 1} // 1
	fmt.Println(singleNumber(nums))
	nums = []int{4, 1, 2, 1, 2} // 4
	fmt.Println(singleNumber(nums))
	nums = []int{1} // 1
	fmt.Println(singleNumber(nums))
}
