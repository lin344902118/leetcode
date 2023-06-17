package main

import "fmt"

/*
921. 使括号有效的最少添加
只有满足下面几点之一，括号字符串才是有效的：
它是一个空字符串，或者
它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
它可以被写作 (A)，其中 A 是有效字符串。
给定一个括号字符串 s ，在每一次操作中，你都可以在字符串的任何位置插入一个括号
例如，如果 s = "()))" ，你可以插入一个开始括号为 "(()))" 或结束括号为 "())))" 。
返回 为使结果字符串 s 有效而必须添加的最少括号数。
*/

func minAddToMakeValid(s string) int {
	if s == "" {
		return 0
	}
	// 左右括号需求量
	leftNeed := 0
	rightNeed := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			rightNeed++
		} else {
			rightNeed--
			// 右括号太多，需要左括号以及重置右括号需求
			if rightNeed == -1 {
				leftNeed++
				rightNeed = 0
			}
		}
	}
	return leftNeed + rightNeed
	// count := 0
	// // 左括号入栈，右括号出栈
	// stack := make([]byte, 0)
	// for i := 0; i < len(s); i++ {
	// 	switch s[i] {
	// 	case '(', '[', '{':
	// 		{
	// 			stack = append(stack, s[i])
	// 		}
	// 	case ')', ']', '}':
	// 		{
	// 			if len(stack) == 0 {
	// 				count++
	// 				continue
	// 			}
	// 			c := stack[len(stack)-1]
	// 			stack = stack[:len(stack)-1]
	// 			if s[i] == ')' && c != '(' {
	// 				// 加上（和另一个
	// 				count += 2
	// 			}
	// 			if s[i] == ']' && c != '[' {
	// 				count += 2
	// 			}
	// 			if s[i] == '}' && c != '{' {
	// 				count += 2
	// 			}
	// 		}
	// 	}
	// }
	// if len(stack) != 0 {
	// 	count += len(stack)
	// }
	// return count
}

func main() {
	s := "())"
	fmt.Println(minAddToMakeValid(s)) // 1
	s = "((("
	fmt.Println(minAddToMakeValid(s)) // 3
}
