package main

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
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		default:
			if len(stack) == 0 {
				return false
			}
			lastEle := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] == ')' && lastEle != '(' {
				return false
			}
			if s[i] == ']' && lastEle != '[' {
				return false
			}
			if s[i] == '}' && lastEle != '{' {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
}
