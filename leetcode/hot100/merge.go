package main

import (
	"fmt"
	"sort"
)

/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	// 先按升序排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] > intervals[j][0] {
			return false
		}
		return true
	})
	n := len(intervals)
	for i := 0; i < n; i++ {
		newStart := intervals[i][0]
		newEnd := intervals[i][1]
		for j := i + 1; j < n; j++ {
			// i包含j
			if intervals[j][0] >= newStart && intervals[j][1] <= newEnd {
				// 跳过j的计算
				i = j
				continue
			}
			// i和j重叠 i在前，j在后
			if intervals[j][0] >= newStart && intervals[j][0] <= newEnd {
				// i和j重叠
				if intervals[j][1] > newEnd {
					newEnd = intervals[j][1]
					// 跳过j的计算
					i = j
					continue
				}
			}
		}
		res = append(res, []int{newStart, newEnd})
	}
	return res
}

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{1, 4},
		{4, 5},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{2, 3},
		{1, 4},
		{3, 5},
		{7, 9},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{1, 4},
		{2, 3},
		{3, 5},
		{7, 9},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{1, 4},
		{0, 1},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{1, 4},
		{2, 3},
	}
	fmt.Println(merge(intervals))
}
