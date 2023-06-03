package main

import "fmt"

/*
  两个数组的交集 II
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，
应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。
可以不考虑输出结果的顺序。
*/

func intersect(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int, 0)
	// 计算nums1中数字出现的次数
	for i := 0; i < len(nums1); i++ {
		cache[nums1[i]]++
	}
	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		count, exist := cache[nums2[i]]
		if exist && count > 0 {
			res = append(res, nums2[i])
			count--
			cache[nums2[i]] = count
		}
	}
	return res
	// sort.Ints(nums1)
	// sort.Ints(nums2)
	// res := make([]int, 0)
	// tmp1, tmp2 := 0, 0
	// for {
	//     if tmp1 >= len(nums1) || tmp2 >= len(nums2) {
	//         break
	//     }
	//     if nums1[tmp1] == nums2[tmp2] {
	//         res = append(res, nums1[tmp1])
	//         tmp1++
	//         tmp2++
	//     } else {
	//         if nums1[tmp1] > nums2[tmp2] {
	//             tmp2++
	//         } else {
	//             tmp1++
	//         }
	//     }
	// }
	// return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2)) // 2 2
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2)) // 9, 4
}
