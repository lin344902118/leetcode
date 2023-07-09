package main

import "fmt"

/*
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	mid := (m + n) / 2
	arr := make([]int, 0, mid+1)
	index1 := 0
	index2 := 0
	for len(arr) != mid+1 {
		// num1全部加入了
		if index1 >= m {
			arr = append(arr, nums2[index2])
			index2++
			// num2全部加入了
		} else if index2 >= n {
			arr = append(arr, nums1[index1])
			index1++
		} else {
			if nums1[index1] < nums2[index2] {
				arr = append(arr, nums1[index1])
				index1++
			} else {
				arr = append(arr, nums2[index2])
				index2++
			}
		}
	}
	// fmt.Println(arr)
	if (m+n)%2 != 0 {
		return float64(arr[mid])
	} else {
		return float64((arr[mid-1] + arr[mid])) / 2.0
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
