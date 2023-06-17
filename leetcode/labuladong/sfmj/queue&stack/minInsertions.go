package main

import "fmt"

/*
1541. 平衡括号字符串的最少插入次数
给你一个括号字符串 s ，它只包含字符 '(' 和 ')' 。一个括号字符串被称为平衡的当它满足：
任何左括号 '(' 必须对应两个连续的右括号 '))' 。
左括号 '(' 必须在对应的连续两个右括号 '))' 之前。
比方说 "())"， "())(())))" 和 "(())())))" 都是平衡的， ")()"， "()))" 和 "(()))" 都是不平衡的。
你可以在任意位置插入字符 '(' 和 ')' 使字符串平衡。
请你返回让 s 平衡的最少插入次数。
*/

func minInsertions(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 左右括号需求量
	leftNeed := 0
	rightNeed := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			rightNeed += 2
			// 对右括号的需求是偶数个，如果是奇数个，需要插入一个左括号
			if rightNeed%2 == 1 {
				leftNeed++
				rightNeed--
			}
		} else {
			rightNeed--
			// 右括号太多，需要左括号以及重置右括号需求
			if rightNeed == -1 {
				leftNeed++
				rightNeed = 1
			}
		}
	}
	return leftNeed + rightNeed
	// // 需要右括号数量
	// count := 0
	// stack := make([]byte, 0, len(s))
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == '(' {
	// 		stack = append(stack, s[i])
	// 	} else {
	// 		if len(stack) == 0 {
	// 			count++
	// 			if i == len(s)-1 {
	// 				count++
	// 				break
	// 			}
	// 			if s[i+1] != ')' {
	// 				count++
	// 			} else {
	// 				i++
	// 			}
	// 			continue
	// 		}
	// 		stack = stack[:len(stack)-1]
	// 		if i == len(s)-1 {
	// 			count++
	// 			break
	// 		}
	// 		if s[i+1] != ')' {
	// 			count++
	// 		} else {
	// 			i++
	// 		}
	// 	}
	// }
	// if len(stack) != 0 {
	// 	count += 2 * len(stack)
	// }
	// return count
}

func main() {
	s := "(()))"
	fmt.Println(minInsertions(s)) // 1
	s = "())"
	fmt.Println(minInsertions(s)) // 0
	s = "))())("
	fmt.Println(minInsertions(s)) // 3
	s = "(()))"                   // 1
	fmt.Println(minInsertions(s))
	s = "(()))(" // 3
	fmt.Println(minInsertions(s))
	s = "(()))((" // 5
	fmt.Println(minInsertions(s))
	s = "(()))(()))" // 2
	fmt.Println(minInsertions(s))
	s = "(()))(()))()" // 3
	fmt.Println(minInsertions(s))
	s = "(()))(()))()())))" // 4
	fmt.Println(minInsertions(s))
}
