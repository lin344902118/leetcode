package main

import "fmt"

/*
315. 计算右侧小于当前元素的个数
给你一个整数数组 nums ，按要求返回一个新数组 counts 。
数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
*/

type Val struct {
	val   int
	index int
}

func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	tmp := make([]Val, len(nums))
	vals := make([]Val, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		vals = append(vals, Val{val: nums[i], index: i})
	}
	merge := func(arr []Val, low, mid, high int) {
		for i := low; i <= high; i++ {
			tmp[i] = arr[i]
		}
		left := low
		right := mid + 1
		for i := low; i <= high; i++ {
			if left == mid+1 {
				arr[i] = tmp[right]
				right++
			} else if right == high+1 {
				arr[i] = tmp[left]
				left++
				res[arr[i].index] += right - mid - 1
			} else if tmp[left].val <= tmp[right].val {
				arr[i] = tmp[left]
				left++
				res[arr[i].index] += right - mid - 1
			} else {
				arr[i] = tmp[right]
				right++
			}
		}
	}
	var sort func([]Val, int, int)
	sort = func(arr []Val, low, high int) {
		if low == high {
			return
		}
		mid := low + (high-low)/2
		sort(arr, low, mid)
		sort(arr, mid+1, high)
		merge(arr, low, mid, high)
		return
	}
	sort(vals, 0, len(nums)-1)
	return res
}

func main() {
	nums := []int{5, 2, 6, 1}
	fmt.Println(countSmaller(nums))
}
