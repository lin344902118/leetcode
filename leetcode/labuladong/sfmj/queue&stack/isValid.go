package main

import "fmt"

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	// 遇到左括号入栈，否则出栈
	tmp := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			{
				tmp = append(tmp, s[i])
			}
		case ')', ']', '}':
			{
				if len(tmp) == 0 {
					return false
				}
				c := tmp[len(tmp)-1]
				tmp = tmp[:len(tmp)-1]
				if s[i] == ')' && c != '(' {
					return false
				}
				if s[i] == ']' && c != '[' {
					return false
				}
				if s[i] == '}' && c != '{' {
					return false
				}
			}
		}
	}
	if len(tmp) != 0 {
		return false
	}
	return true
}

func main() {
	s := "()"
	fmt.Println(isValid(s)) // true
	s = "()[]{}"
	fmt.Println(isValid(s)) // true
	s = "(]"
	fmt.Println(isValid(s)) // false
}
