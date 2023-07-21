package main

/*
72. 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
*/

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	// word1 0-i到word2 0-j的最小编辑距离
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		if word1[i] == word2[j] {
			cache[i][j] = dp(i-1, j-1)
		} else {
			insert := dp(i, j-1) + 1
			replace := dp(i-1, j-1) + 1
			delete := dp(i-1, j) + 1
			cache[i][j] = insert
			if replace < cache[i][j] {
				cache[i][j] = replace
			}
			if delete < cache[i][j] {
				cache[i][j] = delete
			}
		}
		return cache[i][j]
	}
	return dp(m-1, n-1)
}

func main() {
}
