package main

import "fmt"

/*
1170. 比较字符串最小字母出现频次
定义一个函数 f(s)，统计 s  中（按字典序比较）最小字母的出现频次 ，其中 s 是一个非空字符串。

例如，若 s = "dcce"，那么 f(s) = 2，因为字典序最小字母是 "c"，它出现了 2 次。

现在，给你两个字符串数组待查表 queries 和词汇表 words 。对于每次查询 queries[i] ，需统计 words 中满足 f(queries[i]) < f(W) 的 词的数目 ，W 表示词汇表 words 中的每个词。

请你返回一个整数数组 answer 作为答案，其中每个 answer[i] 是第 i 次查询的结果。
*/

func numSmallerByFrequency(queries []string, words []string) []int {
	res := make([]int, 0)
	fWords := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		fWords = append(fWords, f(words[i]))
	}
	for i := 0; i < len(queries); i++ {
		count := 0
		for j := 0; j < len(fWords); j++ {
			if f(queries[i]) < fWords[j] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

func f(s string) int {
	if s == "" {
		return 0
	}
	tmp := make([]int, 26)
	for i := 0; i < len(s); i++ {
		tmp[s[i]-'a']++
	}
	for i := 0; i < len(tmp); i++ {
		if tmp[i] > 0 {
			return tmp[i]
		}
	}
	return 0
}

func main() {
	queries := []string{"cbd"}
	words := []string{"zaaaz"}
	fmt.Println(numSmallerByFrequency(queries, words))
	queries = []string{"bbb", "cc"}
	words = []string{"a", "aa", "aaa", "aaaa"}
	fmt.Println(numSmallerByFrequency(queries, words))
}
