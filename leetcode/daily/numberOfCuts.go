package main

/*
2481. 分割圆的最少切割次数
圆内一个 有效切割 ，符合以下二者之一：

该切割是两个端点在圆上的线段，且该线段经过圆心。
该切割是一端在圆心另一端在圆上的线段。
*/

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n >> 1
	}
	return n
}

func main() {

}
