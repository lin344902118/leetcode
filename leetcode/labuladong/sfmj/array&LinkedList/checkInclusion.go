package main

import "fmt"

/*
567. 字符串的排列
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。
如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。
*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	window := make(map[byte]int, 0)
	need := make(map[byte]int, 0)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		c := s2[right]
		window[c]++
		right++
		count, exist := need[c]
		if exist {
			if count == window[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left == len(s1) {
				return true
			}
			d := s2[left]
			left++
			window[d]--
			count, exist := need[d]
			if exist {
				if window[d] < count {
					valid--
				}
			}
		}
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2)) // true
	s1 = "ab"
	s2 = "eidboaoo"
	fmt.Println(checkInclusion(s1, s2)) // false
}
