package main

import (
	"fmt"
)

/*
752. 打开转盘锁
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字：
'0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，
这个锁将会被永久锁定，无法再被旋转。

字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
*/

func plusOne(s string, j int) string {
	bs := []byte(s)
	if bs[j] == '9' {
		bs[j] = '0'
	} else {
		bs[j]++
	}
	return string(bs)
}

func minusOne(s string, j int) string {
	bs := []byte(s)
	if bs[j] == '0' {
		bs[j] = '9'
	} else {
		bs[j]--
	}
	return string(bs)
}

func openLock(deadends []string, target string) int {
	queue := make([]string, 0)
	visited := make(map[string]struct{}, 0)
	deads := make(map[string]struct{}, 0)
	queue = append(queue, "0000")
	visited["0000"] = struct{}{}
	// 将死锁加入deads列表中
	for _, d := range deadends {
		deads[d] = struct{}{}
	}
	step := 0
	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			_, e := deads[cur]
			if e {
				continue
			}
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				_, e := visited[up]
				if !e {
					queue = append(queue, up)
					visited[up] = struct{}{}
				}
				down := minusOne(cur, j)
				_, e = visited[down]
				if !e {
					queue = append(queue, down)
					visited[down] = struct{}{}
				}
			}
		}
		step++
	}
	return -1
}

func main() {
	deadends := []string{
		"0201", "0101", "0102", "1212", "2002",
	}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}
