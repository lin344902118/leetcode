package main

import "fmt"

/*
  计数质数
给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
*/

func countPrimes(n int) int {
	arr := make([]bool, n)
	cnt := 0
	for i := 2; i < n; i++ {
		if arr[i] {
			continue
		}
		cnt++
		for j := i; j < n; j += i {
			arr[j] = true
		}
	}
	return cnt
}

func main() {
	n := 10
	fmt.Println(countPrimes(n))
}
