package main

import (
	"fmt"
	"sort"
)

/*
  老鼠和奶酪
有两只老鼠和 n 块不同类型的奶酪，每块奶酪都只能被其中一只老鼠吃掉。

下标为 i 处的奶酪被吃掉的得分为：

如果第一只老鼠吃掉，则得分为 reward1[i] 。
如果第二只老鼠吃掉，则得分为 reward2[i] 。
给你一个正整数数组 reward1 ，一个正整数数组 reward2 ，和一个非负整数 k 。

请你返回第一只老鼠恰好吃掉 k 块奶酪的情况下，最大 得分为多少。
*/

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	loss := make([]int, 0, len(reward1))
	score := 0
	for i := 0; i < len(reward1); i++ {
		loss = append(loss, reward1[i]-reward2[i])
		score += reward2[i]
	}
	sort.Ints(loss)
	// 题目保证k小于reward1长度，所以不会越界
	for i := len(loss) - 1; i >= len(loss)-k; i-- {
		score += loss[i]
	}
	return score
}

func main() {
	reward1 := []int{1, 1, 3, 4}
	reward2 := []int{4, 4, 1, 1}
	k := 2
	fmt.Println(miceAndCheese(reward1, reward2, k))
	fmt.Println(miceAndCheese(reward1, reward2, 3))
	reward1 = []int{1, 1}
	reward2 = []int{1, 1}
	fmt.Println(miceAndCheese(reward1, reward2, k))
}
