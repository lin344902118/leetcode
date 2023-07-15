package main

import "fmt"

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

var cache = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		letters := cache[digits[index]]
		for _, letter := range letters {
			path += letter
			backtrack(index+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, "")
	return res
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
	digits = ""
	fmt.Println(letterCombinations(digits))
	digits = "2"
	fmt.Println(letterCombinations(digits))
}
