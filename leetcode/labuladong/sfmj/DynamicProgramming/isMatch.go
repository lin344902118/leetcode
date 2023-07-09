package main

/*
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
*/

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	var dp func(s, p string, i, j int) bool
	dp = func(s, p string, i, j int) bool {
		if j == n {
			return i == m
		}
		if i == m {
			if (n-j)%2 == 1 {
				return false
			}
			for ; j+1 < n; j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}
		if cache[i][j] != -1 {
			if cache[i][j] == 0 {
				return false
			} else {
				return true
			}
		}
		res := false
		if s[i] == p[j] || p[j] == '.' {
			if j < n-1 && p[j+1] == '*' {
				res = dp(s, p, i, j+2) || dp(s, p, i+1, j)
			} else {
				res = dp(s, p, i+1, j+1)
			}
		} else {
			if j < n-1 && p[j+1] == '*' {
				res = dp(s, p, i, j+2)
			} else {
				res = false
			}
		}
		if res {
			cache[i][j] = 1
		} else {
			cache[i][j] = 0
		}
		return res
	}
	return dp(s, p, 0, 0)
}

func main() {
}
