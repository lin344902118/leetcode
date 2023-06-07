package main

/*
  最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/

func longestCommonPrefix(strs []string) string {
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(s) {
			s = strs[i]
		}
	}
	prefix := make([]byte, 0, len(s))
	for j := 0; j < len(s); j++ {
		same := true
		for k := 0; k < len(strs); k++ {
			if strs[k][j] != s[j] {
				same = false
				break
			}
		}
		if same {
			prefix = append(prefix, s[j])
		} else {
			break
		}
	}
	return string(prefix)
}

func main() {

}
