package main

import (
	"fmt"
)

/*
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int, 0)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		window[c]++
		right++
		for window[c] > 1 {
			d := s[left]
			window[d]--
			left++
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s)) // 3
	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s)) // 1
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s)) // 3
	s = ""
	fmt.Println(lengthOfLongestSubstring(s)) // 0
}
