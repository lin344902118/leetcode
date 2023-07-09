package main

import (
	"fmt"
	"sort"
)

/*
354. 俄罗斯套娃信封问题
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。
*/

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			if envelopes[i][1] > envelopes[j][1] {
				return true
			}
		}
		return false
	})
	tmp := make([]int, len(envelopes))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = envelopes[i][1]
	}
	return lengthOfLIS(tmp)
}

func lengthOfLIS(nums []int) int {
	top := make([]int, len(nums))
	piles := 0
	for i := 0; i < len(nums); i++ {
		poker := nums[i]
		left := 0
		right := piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

func main() {
	env := [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}
	fmt.Println(maxEnvelopes(env))
}
