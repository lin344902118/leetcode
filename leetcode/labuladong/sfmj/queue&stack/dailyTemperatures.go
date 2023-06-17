package main

import "fmt"

/*
739. 每日温度
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
如果气温在这之后都不会升高，请在该位置用 0 来代替。
*/

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 单调栈,记录索引，而非值
	stack := make([]int, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		num := temperatures[i]
		for len(stack) > 0 && num >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// 堆栈是倒叙的，所以答案要反转一下
		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = stack[len(stack)-1] - i
		}
		// 记录索引
		stack = append(stack, i)
	}
	return res
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	// fmt.Println(dailyTemperatures(temperatures))
	temperatures = []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	fmt.Println(dailyTemperatures(temperatures))
}
