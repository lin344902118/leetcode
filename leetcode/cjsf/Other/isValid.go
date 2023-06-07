package main

import "fmt"

/*
  有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func isValid(s string) bool {
	// 合法的括号长度必须是偶数
	if len(s)%2 != 0 {
		return false
	}
	valid := true
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 {
				valid = false
				break
			}
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] == ')' && tmp != '(' {
				valid = false
				break
			}
			if s[i] == ']' && tmp != '[' {
				valid = false
				break
			}
			if s[i] == '}' && tmp != '{' {
				valid = false
				break
			}
		}
	}
	if len(stack) != 0 {
		valid = false
	}
	return valid
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
	s = "()[]{}"
	fmt.Println(isValid(s))
	s = "(]"
	fmt.Println(isValid(s))
}
