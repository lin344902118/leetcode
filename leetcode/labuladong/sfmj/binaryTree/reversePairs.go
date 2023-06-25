package main

import "fmt"

/*
493. 翻转对
给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
你需要返回给定数组中的重要翻转对的数量。
*/
func reversePairs(nums []int) int {
	res := 0
	// 归并排序
	tmp := make([]int, len(nums))
	merge := func(arr []int, low, mid, high int) {
		for i := low; i <= high; i++ {
			tmp[i] = arr[i]
		}
		end := mid + 1
		for i := low; i <= mid; i++ {
			for end <= high && (int64(nums[i]) > int64(nums[end]*2)) {
				end++
			}
			res += end - mid - 1
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
	return res
}

func main() {
	nums := []int{1, 3, 2, 3, 1}
	nums = []int{2, 4, 3, 5, 1}
	fmt.Println(reversePairs(nums))
}
