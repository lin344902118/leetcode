package main

import (
	"fmt"
)

/*
139. 单词拆分
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。
请你判断是否可以利用字典中出现的单词拼接出 s 。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
*/

func wordBreak(s string, wordDict []string) bool {
	m := len(s)
	cache := make([]int, m)
	for i := 0; i < m; i++ {
		cache[i] = -1
	}
	var dp func(index int) bool
	dp = func(index int) bool {
		if index == m {
			return true
		}
		if cache[index] != -1 {
			if cache[index] == 1 {
				return true
			} else {
				return false
			}
		}
		for _, word := range wordDict {
			length := len(word)
			if s[index] != word[0] || index+length > m || s[index:index+length] != word {
				continue
			}
			if dp(index + length) {
				cache[index] = 1
				return true
			}
		}
		cache[index] = 0
		return false
	}
	return dp(0)
}

func main() {
	s := "leetcode"
	wordDict := []string{
		"leet",
		"code",
	}
	fmt.Println(wordBreak(s, wordDict)) // true
	s = "applepenapple"
	wordDict = []string{
		"apple",
		"pen",
	}
	fmt.Println(wordBreak(s, wordDict))
	s = "catsandog"
	wordDict = []string{
		"cats",
		"dog",
		"sand",
		"and",
		"cat",
	}
	fmt.Println(wordBreak(s, wordDict))
}
