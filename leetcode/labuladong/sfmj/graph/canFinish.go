package main

import "fmt"

/*
207. 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，
其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
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
		onPath[s] = false
	}
	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}
	return !hasCycle
}

func buildGraph(num int, requires [][]int) [][]int {
	graph := make([][]int, num)
	for i := 0; i < len(requires); i++ {
		// requires[i]长度固定为2
		graph[requires[i][1]] = append(graph[requires[i][1]], requires[i][0])
	}
	return graph
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses, prerequisites))
}
