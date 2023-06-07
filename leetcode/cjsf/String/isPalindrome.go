package main

import "strings"

/*
  验证回文串
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
*/

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for {
		if left >= right {
			break
		}
		if !isLetterOrNumber(s[left]) {
			left++
			continue
		}
		if !isLetterOrNumber(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetterOrNumber(s byte) bool {
	if (s >= '0' && s <= '9') || (s >= 'a' && s <= 'z') {
		return true
	}
	return false
}

func main() {

}
