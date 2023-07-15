package main

import "fmt"

/*
32. 最长有效括号
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
*/

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	// dp[i] 的定义：记录以 s[i-1] 结尾的最长合法括号子串长度
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, i)
			dp[i+1] = 0
		case ')':
			if len(stack) > 0 {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				dp[i+1] = i + 1 - last + dp[last]
			} else {
				dp[i+1] = 0
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func main() {
	s := "(()"
	fmt.Println(longestValidParentheses(s))
	s = ")()())"
	fmt.Println(longestValidParentheses(s))
	s = ""
	fmt.Println(longestValidParentheses(s))
	s = "()(()"
	fmt.Println(longestValidParentheses(s))
}
