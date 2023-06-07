package main

import (
	"fmt"
	"strconv"
)

/*
  汉明距离
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
*/

func hammingDistance(x int, y int) int {
	tmp1 := strconv.FormatInt(int64(x), 2)
	tmp2 := strconv.FormatInt(int64(y), 2)
	loss := 0
	if len(tmp2)-len(tmp1) > 0 {
		loss = len(tmp2) - len(tmp1)
		s := ""
		for i := 0; i < loss; i++ {
			s += "0"
		}
		tmp1 = s + tmp1
	} else {
		loss = len(tmp1) - len(tmp2)
		s := ""
		for i := 0; i < loss; i++ {
			s += "0"
		}
		tmp2 = s + tmp2
	}
	res := 0
	for i := 0; i < len(tmp1); i++ {
		if tmp1[i] != tmp2[i] {
			res++
		}
	}
	return res
}

func main() {
	x := 1
	y := 4
	fmt.Println(hammingDistance(x, y))
}
