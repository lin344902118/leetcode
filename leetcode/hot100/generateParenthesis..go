package main

import "fmt"

/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/

func generateParenthesis(n int) []string {
	// // 思路一
	// leftCount := n
	// rightCount := n
	// res := make([]string, 0)
	// path := ""
	// stack := make([]string, 0)
	// var backtrack func(path string, index int)
	// backtrack = func(path string, index int) {
	// 	if index == n*2 {
	// 		res = append(res, path)
	// 		return
	// 	}
	// 	if len(stack) == 0 {
	// 		path += "("
	// 		leftCount--
	// 		stack = append(stack, "(")
	// 		backtrack(path, index+1)
	// 		path = path[:len(path)-1]
	// 		leftCount++
	// 		stack = stack[:len(stack)-1]
	// 		return
	// 	}
	// 	for _, p := range []string{"(", ")"} {
	// 		if p == "(" && leftCount == 0 {
	// 			continue
	// 		}
	// 		if p == ")" && rightCount == 0 {
	// 			continue
	// 		}
	// 		path += p
	// 		if p == "(" {
	// 			leftCount--
	// 			stack = append(stack, p)
	// 		} else {
	// 			rightCount--
	// 			stack = stack[:len(stack)-1]
	// 		}
	// 		backtrack(path, index+1)
	// 		path = path[:len(path)-1]
	// 		if p == "(" {
	// 			leftCount++
	// 			stack = stack[:len(stack)-1]
	// 		} else {
	// 			rightCount++
	// 			stack = append(stack, p)
	// 		}
	// 	}
	// }
	// backtrack(path, 0)
	// return res
	// 思路二
	leftCount := 0
	rightCount := 0
	res := make([]string, 0)
	path := ""
	var backtrack func(path string, index int)
	backtrack = func(path string, index int) {
		if len(path) == n*2 {
			res = append(res, path)
			return
		}
		if leftCount < n {
			path += "("
			leftCount++
			backtrack(path, index+1)
			path = path[:len(path)-1]
			leftCount--
		}
		if rightCount < leftCount {
			path += ")"
			rightCount++
			backtrack(path, index+1)
			path = path[:len(path)-1]
			rightCount--
		}
	}
	backtrack(path, 0)
	return res
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
