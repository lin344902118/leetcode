package main

/*
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func lengthOfLongestSubstring(s string) int {
	// 思路一
	// if len(s) == 0 {
	// 	return 0
	// }
	// max := 0
	// for i := 0; i < len(s); i++ {
	// 	m := make(map[byte]struct{}, 0)
	// 	m[s[i]] = struct{}{}
	// 	length := 1
	// 	for j := i + 1; j < len(s); j++ {
	// 		_, e := m[s[j]]
	// 		if e {
	// 			break
	// 		}
	// 		m[s[j]] = struct{}{}
	// 		length++
	// 	}
	// 	if length > max {
	// 		max = length
	// 	}
	// }
	// return max
	// 思路二
	window := make(map[byte]int, 0)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		// right会加一,所以不能直接使用s[right]
		c := s[right]
		// 调换right和window[c]顺序会变慢,不知道为啥
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

func main() {
}
