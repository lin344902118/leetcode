package main

import (
	"fmt"
	"sort"
)

/*
49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
*/

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	// 用map缓存，key为按字典序排序后得到的值，value是字母异位词组合
	cache := make(map[string][]string, 0)
	for _, str := range strs {
		tmp := make([]string, 0, len(str))
		for i := 0; i < len(str); i++ {
			tmp = append(tmp, string(str[i]))
		}
		sort.Strings(tmp)
		key := ""
		for i := 0; i < len(tmp); i++ {
			key += tmp[i]
		}
		_, e := cache[key]
		if !e {
			cache[key] = make([]string, 0)
		}
		cache[key] = append(cache[key], str)
	}
	for _, v := range cache {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
}
