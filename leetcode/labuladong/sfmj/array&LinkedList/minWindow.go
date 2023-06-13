package main

import "fmt"

/*
76. 最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	need := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	left, right := 0, 0
	valid := 0
	start, end := 0, len(s)+1
	for i := 0; i < len(t); i++ {
		need[t[i]]++
		window[t[i]] = 0
	}
	for right < len(s) {
		c := s[right]
		window[c]++
		right++
		count, exist := need[c]
		if exist {
			if window[c] == count {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < end-start {
				start = left
				end = right
			}
			d := s[left]
			left++
			window[d]--
			_, exist := need[d]
			if exist {
				if window[d] < need[d] {
					valid--
				}
			}
		}
	}
	if end == len(s)+1 {
		return ""
	}
	return s[start:end]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	// fmt.Println(minWindow(s, t))
	// s = "a"
	// t = "a"
	// fmt.Println(minWindow(s, t))
	// s = "a"
	// t = "b"
	// fmt.Println(minWindow(s, t))
	s = "cabwefgewcwaefgcf"
	t = "cae"
	fmt.Println(minWindow(s, t))
}
