package main

import (
	"fmt"
	"strings"
)

/*
316. 去除重复字母
给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。
需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
*/

func removeDuplicateLetters(s string) string {
	// 单调栈，确保字典序最小
	stack := make([]byte, 0)
	// 记录每个字母出现次数，确保字符相对位置
	count := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	// 判断字母是否在栈中，确保唯一
	inStack := make([]bool, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']--
		if inStack[s[i]-'a'] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > s[i] {
			if count[stack[len(stack)-1]-'a'] == 0 {
				break
			}
			e := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inStack[e-'a'] = false
		}
		stack = append(stack, s[i])
		inStack[s[i]-'a'] = true
	}
	sb := strings.Builder{}
	for i := 0; i < len(stack); i++ {
		sb.WriteByte(stack[i])
	}
	return sb.String()
}

func main() {
	s := "bcabc"
	// fmt.Println(removeDuplicateLetters(s))
	s = "cbacdcbc"
	fmt.Println(removeDuplicateLetters(s))
}
