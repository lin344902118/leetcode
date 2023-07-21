package main

import (
	"fmt"
)

/*
75. 颜色分类
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，
原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。
*/

func sortColors(nums []int) {
	// 思路一
	// // 用一个长度为3的[]int记录红色、白色、蓝色出现次数，然后再换
	// counts := make([]int, 3)
	// for _, num := range nums {
	// 	counts[num]++
	// }
	// count := 0
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < counts[i]; j++ {
	// 		nums[count+j] = i
	// 	}
	// 	count += counts[i]
	// }
	// return
	// 思路二
	// 两次遍历，第一次将0交换到最前面，第二次交换1
	index := 0
	m := len(nums)
	for i := 0; i < m; i++ {
		if nums[i] == 0 {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	for i := index; i < m; i++ {
		if nums[i] == 1 {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
}

func main() {
	nums := []int{
		2, 0, 2, 1, 1, 0,
	}
	sortColors(nums)
	fmt.Println(nums)
}
