package main

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
*/

func longestPalindrome(s string) string {
	// 思路一
	// if len(s) == 0 {
	// 	return ""
	// }
	// res := ""
	// for i := 0; i < len(s); i++ {
	// 	for j := i + 1; j <= len(s); j++ {
	// 		if isPalindrome(s[i:j]) {
	// 			if (j - i) > len(res) {
	// 				res = s[i:j]
	// 			}
	// 		}
	// 	}
	// }
	// return res
	// 思路二
	res := ""
	for i := 0; i < len(s); i++ {
		tmp := ""
		tmp1 := palindrome(s, i, i)
		tmp2 := palindrome(s, i, i+1)
		if len(tmp1) > len(tmp2) {
			tmp = tmp1
		} else {
			tmp = tmp2
		}
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func palindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 当前left和right的前一步才是回文字符串
	return s[left+1 : right]
}

func main() {
}
