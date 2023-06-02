package main

import "fmt"

/*
 删除数组中的偶数元素，要求时间复杂度O(n),空间复杂度O(1)
*/
func removeEvenNumber(arr []int) []int {
	if arr == nil {
		return arr
	}
	// 左指针找偶数，右指针找奇数
	left := 0
	right := len(arr) - 1
	// 直到两个指针相逢
	for left < right {
		// 左边找到第一个偶数
		for left < len(arr) && arr[left]%2 != 0 {
			left++
		}
		// 右边找到第一个奇数
		for right >= 0 && arr[right]%2 == 0 {
			right--
		}
		// 左右交换
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	return arr[:left]
}

func main() {
	arr := []int{1, 2, 3} // 1, 3
	fmt.Println(removeEvenNumber(arr))
	arr = []int{1, 1, 2, 3} // 1, 1, 3
	fmt.Println(removeEvenNumber(arr))
	arr = []int{1, 1, 2, 2, 3} // 1, 1, 3
	fmt.Println(removeEvenNumber(arr))
	arr = []int{1, 1, 2, 2, 3, 3} // 1, 1, 3, 3
	fmt.Println(removeEvenNumber(arr))
	arr = []int{1, 3, 5}
	fmt.Println(removeEvenNumber(arr)) // 1,3,5
	arr = []int{2, 4, 6}
	fmt.Println(removeEvenNumber(arr)) //
	arr = []int{1, 2, 3, 4}
	fmt.Println(removeEvenNumber(arr)) // 1,3
	arr = []int{2, 3, 4}
	fmt.Println(removeEvenNumber(arr)) // 3
	arr = []int{2, 3, 4, 5}
	fmt.Println(removeEvenNumber(arr)) // 5,3
}
