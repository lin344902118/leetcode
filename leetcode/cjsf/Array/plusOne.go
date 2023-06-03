package main

import (
	"fmt"
)

/*
  加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits
	// // digits加一后可能进一位所以长度比digits大一位
	// res := make([]int, len(digits)+1)
	// flag := true
	// for i := len(digits) - 1; i >= 0; i-- {
	// 	if digits[i] == 9 && flag {
	// 		res[i+1] = 0
	// 		flag = true
	// 	} else {
	// 		if flag {
	// 			res[i+1] = digits[i] + 1
	// 			flag = false
	// 		} else {
	// 			res[i+1] = digits[i]
	// 		}
	// 	}
	// }
	// if flag {
	// 	res[0] = 1
	// 	return res
	// } else {
	// 	return res[1:]
	// }
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits)) // 1 2 4
	digits = []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits)) // 4 3 2 2
	digits = []int{0}
	fmt.Println(plusOne(digits)) // 1
	digits = []int{9, 9, 9}
	fmt.Println(plusOne(digits)) // 1 0 0 0
}
