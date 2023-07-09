package main

import (
	"fmt"
	"math"
)

/*
514. 自由之路
电子游戏“辐射4”中，任务 “通向自由” 要求玩家到达名为 “Freedom Trail Ring” 的金属表盘，
并使用表盘拼写特定关键词才能开门。

给定一个字符串 ring ，表示刻在外环上的编码；给定另一个字符串 key ，
表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。

最初，ring 的第一个字符与 12:00 方向对齐。您需要顺时针或逆时针旋转 ring
以使 key 的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。

旋转 ring 拼出 key 字符 key[i] 的阶段中：

您可以将 ring 顺时针或逆时针旋转 一个位置 ，计为1步。
旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于字符 key[i] 。
如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。
按完之后，您可以开始拼写 key 的下一个字符（下一阶段）, 直至完成所有拼写。
*/

func findRotateSteps(ring string, key string) int {
	m := len(ring)
	n := len(key)
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	char2Index := make(map[byte][]int, 0)
	for i := 0; i < m; i++ {
		_, exist := char2Index[ring[i]]
		if !exist {
			char2Index[ring[i]] = make([]int, 0)
		}
		char2Index[ring[i]] = append(char2Index[ring[i]], i)
	}
	// dp表示从ring i到key j需要操作的次数
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if j == len(key) {
			return 0
		}
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		res := math.MaxInt32
		for _, k := range char2Index[key[j]] {
			var delta int
			if k-i < 0 {
				delta = i - k
			} else {
				delta = k - i
			}
			if m-delta < delta {
				delta = m - delta
			}
			subProblem := dp(k, j+1)
			if 1+delta+subProblem < res {
				res = 1 + delta + subProblem
			}
		}
		cache[i][j] = res
		return res
	}
	return dp(0, 0)
}

func main() {
	ring := "godding"
	key := "gd"
	fmt.Println(findRotateSteps(ring, key))
	ring = "abcde"
	key = "ade"
	// fmt.Println(findRotateSteps(ring, key))
	ring = "pqwcx"
	key = "cpqwx"
	// fmt.Println(findRotateSteps(ring, key))
	ring = "iotfo"
	key = "fioot"
	// fmt.Println(findRotateSteps(ring, key))
}
