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
	m := len(s)
	n := len(t)
	if n > m {
		return ""
	}
	// 滑动窗口
	need := make(map[byte]int, 0)
	for i := 0; i < n; i++ {
		need[t[i]]++
	}
	left := 0
	right := 0
	valid := 0
	window := make(map[byte]int, 0)
	start, end := 0, m+1
	for right < m {
		// 扩大窗口
		c := s[right]
		window[c]++
		right++
		count, e := need[c]
		if e {
			if window[c] == count {
				valid++
			}
		}
		// 所有条件都满足了，缩小窗口
		for valid == len(need) {
			if right-left < end-start {
				start = left
				end = right
			}
			d := s[left]
			left++
			window[d]--
			_, e := need[d]
			if e {
				if window[d] < need[d] {
					valid--
				}
			}
		}
	}
	if end == m+1 {
		return ""
	}
	return s[start:end]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	s = "a"
	t = "b"
	fmt.Println(minWindow(s, t))
}
