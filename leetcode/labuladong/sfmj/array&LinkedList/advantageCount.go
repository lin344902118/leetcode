package main

import (
	"container/heap"
	"sort"
)

/*
870. 优势洗牌
给定两个长度相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足
nums1[i] > nums2[i] 的索引 i 的数目来描述。
返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。
*/

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][1] > pq[j][1]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	// 给 nums2 降序排序
	maxPq := make(PriorityQueue, 0)
	heap.Init(&maxPq)

	for i := 0; i < n; i++ {
		pair := []int{i, nums2[i]}
		heap.Push(&maxPq, pair)
	}

	// 给 nums1 升序排序
	sort.Ints(nums1)

	// nums1[left] 是最小值，nums1[right] 是最大值
	left, right := 0, n-1
	res := make([]int, n)

	for maxPq.Len() > 0 {
		pair := heap.Pop(&maxPq).([]int)
		// maxval 是 nums2 中的最大值，i 是对应索引
		i, maxval := pair[0], pair[1]
		if maxval < nums1[right] {
			// 如果 nums1[right] 能胜过 maxval，那就自己上
			res[i] = nums1[right]
			right--
		} else {
			// 否则用最小值混一下，养精蓄锐
			res[i] = nums1[left]
			left++
		}
	}
	return res

}

func main() {

}
