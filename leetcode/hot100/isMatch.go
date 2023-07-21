package main

/*
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
*/

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	// 使用缓存加快匹配速度
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	// 判断s中i之后的是否能被p中j之后匹配
	var match func(i, j int) bool
	match = func(i, j int) bool {
		// p已经被匹配完了
		if j == n {
			// 判断s是否也是匹配完
			return i == m
		}
		// s匹配完了，p还没匹配完。可能p后面是一堆的x*
		if i == m {
			// 剩余奇数个，不可能是x*
			if (n-j)%2 == 1 {
				return false
			}
			// p后面的都是x*，因为*可以不匹配，所以这些可以忽略
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
		// 判断s[i]是否等于p[j]或者p[j]等于.
		if s[i] == p[j] || p[j] == '.' {
			// p[j+1]等于*,可以匹配也可以不匹配
			if j < n-1 && p[j+1] == '*' {
				// match(i, j+2)表示不匹配i，match(i+1,j)表示匹配i
				res = match(i, j+2) || match(i+1, j)
			} else {
				// 匹配下一个
				res = match(i+1, j+1)
			}
		} else {
			// s[i]不等于p[j]只能不匹配
			if j < n-1 && p[j+1] == '*' {
				res = match(i, j+2)
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
	return match(0, 0)
}

func main() {
}
