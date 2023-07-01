package main

import "fmt"

/*
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/

func combinationSum3(k int, n int) [][]int {

	res := [][]int{}
	path := []int{}

	var backtrace func(currNum, leftSum int)
	backtrace = func(currNum, leftSum int) {
		if len(path) == k && leftSum == 0 {
			res = append(res, append([]int{}, path...))
			return
		}

		if leftSum < 0 || len(path)+10-currNum < k {
			return
		}

		backtrace(currNum+1, leftSum)

		path = append(path, currNum)
		backtrace(currNum+1, leftSum-currNum)
		path = path[:len(path)-1]
	}

	backtrace(1, n)

	return res

}

// func combinationSum3(k int, n int) [][]int {
// 	res := make([][]int, 0)
// 	if n > 45 || k > 10 {
// 		return res
// 	}
// 	used := make([]bool, 10)
// 	var backtrack func(path []int, num, sum, k, target int, used []bool)
// 	backtrack = func(path []int, num, sum, k, target int, used []bool) {
// 		if k == 0 {
// 			if sum == target {
// 				res = append(res, append([]int{}, path...))
// 			}
// 			return
// 		}
// 		for i := num + 1; i < 10; i++ {
// 			if used[i] {
// 				continue
// 			}
// 			if sum+i > target {
// 				break
// 			}
// 			sum += i
// 			used[i] = true
// 			path = append(path, i)
// 			fmt.Println(sum, used, path)
// 			backtrack(path, i, sum, k-1, target, used)
// 			sum -= i
// 			used[i] = false
// 			path = path[:len(path)-1]
// 		}
// 		return
// 	}
// 	backtrack([]int{}, 0, 0, k, n, used)
// 	return res
// }

func main() {
	// fmt.Println(combinationSum3(3, 7))
	// fmt.Println(combinationSum3(3, 9))
	// fmt.Println(combinationSum3(4, 1))
	fmt.Println(combinationSum3(4, 24))
}
