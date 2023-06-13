package main

import "fmt"

/*
438. 找到字符串中所有字母异位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
*/

func findAnagrams(s string, p string) []int {
	need := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right, valid := 0, 0, 0
	res := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		count, exist := need[c]
		if exist {
			if window[c] == count {
				valid++
			}
		}
		for valid == len(need) {
			if right-left == len(p) {
				res = append(res, left)
			}
			d := s[left]
			window[d]--
			left++
			count, exist := need[d]
			if exist {
				if window[d] < count {
					valid--
				}
			}
		}
	}
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p)) // [0,6]
	s = "abab"
	p = "ab"
	fmt.Println(findAnagrams(s, p)) // [0,1,2]
}
