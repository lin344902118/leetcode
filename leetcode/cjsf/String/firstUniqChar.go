package main

/*
  字符串中的第一个唯一字符
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
*/

func firstUniqChar(s string) int {
	arr := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		arr[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		v := arr[s[i]]
		if v == 1 {
			return i
		}
	}
	return -1
}

func main() {

}
