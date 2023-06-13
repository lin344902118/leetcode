package main

/*
1094. 拼车
车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）

给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi]
表示第 i 次旅行有 numPassengersi 乘客，
接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。

当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
*/

func carPooling(trips [][]int, capacity int) bool {
	tmp := make([]int, 1001)
	for i := 0; i < len(trips); i++ {
		tmp[trips[i][1]] += trips[i][0]
		tmp[trips[i][2]] -= trips[i][0]
	}
	sum := 0
	for i := 0; i < len(tmp); i++ {
		sum += tmp[i]
		if sum > capacity {
			return false
		}
	}
	return true
	// nums := make([]int, 1000)
	// for _, trip := range trips {
	// 	for i := trip[1]; i < trip[2]; i++ {
	// 		nums[i] += trip[0]
	// 		if nums[i] > capacity {
	// 			return false
	// 		}
	// 	}
	// }
	// return true
}

func main() {

}
