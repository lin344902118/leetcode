package main

import "fmt"

/*
2178. 拆分成最多数目的正偶数之和
给你一个整数 finalSum 。请你将它拆分成若干个 互不相同 的正偶数之和，且拆分出来的正偶数数目 最多 。

比方说，给你 finalSum = 12 ，那么这些拆分是 符合要求 的（互不相同的正偶数且和为 finalSum）：
(2 + 10) ，(2 + 4 + 6) 和 (4 + 8) 。它们中，(2 + 4 + 6) 包含最多数目的整数。注意 finalSum 不能拆分成 (2 + 2 + 4 + 4) ，因为拆分出来的整数必须互不相同。
请你返回一个整数数组，表示将整数拆分成 最多 数目的正偶数数组。
如果没有办法将 finalSum 进行拆分，请你返回一个 空 数组。你可以按 任意 顺序返回这些整数。
*/

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	res := make([]int64, 0)
	var i int64
	for i = 2; i <= finalSum; i += 2 {
		res = append(res, i)
		finalSum -= i
	}
	res[len(res)-1] += finalSum
	return res
}

func main() {
	var num int64 = 8
	fmt.Println(maximumEvenSplit(num))
}
