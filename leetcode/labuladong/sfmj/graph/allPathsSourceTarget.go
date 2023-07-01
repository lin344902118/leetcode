package main

import "fmt"

/*
797. 所有可能的路径
给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。
*/

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	var traverse func([][]int, int, []int)
	traverse = func(graph [][]int, start int, path []int) {
		path = append(path, start)
		if start == len(graph)-1 {
			// path是会变的，需要复制一份
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(graph[start]); i++ {
			traverse(graph, graph[start][i], path)
		}
		path = path[:len(path)-1]
	}
	traverse(graph, 0, []int{})
	return res
}

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	graph = [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}
