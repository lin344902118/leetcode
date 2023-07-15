package main

import (
	"fmt"
	"strconv"
)

/*
338. 比特位计数
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，
返回一个长度为 n + 1 的数组 ans 作为答案。
*/

func countBits(n int) []int {
	res := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		tmp := strconv.FormatInt(int64(i), 2)
		count := 0
		for j := 0; j < len(tmp); j++ {
			if tmp[j] == '1' {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

func main() {
	n := 5
	fmt.Println(countBits(n))
}
