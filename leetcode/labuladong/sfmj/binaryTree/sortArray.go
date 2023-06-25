package main

import "fmt"

/*
912. 排序数组
给你一个整数数组 nums，请你将该数组升序排列。
*/

func sortArray(nums []int) []int {
	// 归并排序
	tmp := make([]int, len(nums))
	merge := func(arr []int, low, mid, high int) {
		for i := low; i <= high; i++ {
			tmp[i] = arr[i]
		}
		left := low
		right := mid + 1
		for i := low; i <= high; i++ {
			// 左数组没有元素了
			if left == mid+1 {
				arr[i] = tmp[right]
				right++
				continue
			} else if right == high+1 {
				// 右数组没有元素了
				arr[i] = tmp[left]
				left++
				continue
			}
			if tmp[left] <= tmp[right] {
				// 取左右数组中最小值
				arr[i] = tmp[left]
				left++
			} else {
				arr[i] = tmp[right]
				right++
			}
		}
	}
	var sort func([]int, int, int)
	sort = func(arr []int, low, high int) {
		if low == high {
			return
		}
		mid := low + (high-low)/2
		sort(arr, low, mid)
		sort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}
