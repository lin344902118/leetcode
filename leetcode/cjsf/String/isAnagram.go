package main

/*
  有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sArr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		sArr[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		if sArr[t[i]-'a'] == 0 {
			return false
		}
		sArr[t[i]-'a']--
	}
	return true
}

func main() {

}
