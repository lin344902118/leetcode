package main

import "fmt"

/*
  3的幂
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
*/

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for {
		if n < 3 {
			return false
		}
		if n == 3 {
			return true
		}
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
}

func main() {
	n := 27
	fmt.Println(isPowerOfThree(n))
	n = 0
	fmt.Println(isPowerOfThree(n))
	n = 9
	fmt.Println(isPowerOfThree(n))
	n = 45
	fmt.Println(isPowerOfThree(n))
}
