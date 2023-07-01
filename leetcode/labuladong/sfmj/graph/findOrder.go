package main

import "fmt"

/*
210. 课程表 II
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，
其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。
如果不可能完成所有课程，返回 一个空数组 。
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0)
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false
	var traverse func([][]int, int)
	traverse = func(graph [][]int, s int) {
		if onPath[s] {
			hasCycle = true
		}
		if visited[s] || hasCycle {
			return
		}
		visited[s] = true
		onPath[s] = true
		for i := 0; i < len(graph[s]); i++ {
			traverse(graph, graph[s][i])
		}
		res = append(res, s)
		onPath[s] = false
	}
	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}
	if hasCycle {
		return []int{}
	}
	return res
}

func buildGraph(num int, requires [][]int) [][]int {
	graph := make([][]int, num)
	for i := 0; i < len(requires); i++ {
		// requires[i]长度固定为2
		graph[requires[i][0]] = append(graph[requires[i][0]], requires[i][1])
	}
	return graph
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}
	fmt.Println(findOrder(numCourses, prerequisites))
}
