package main

import "fmt"

/*
1011. 在 D 天内送达包裹的能力
传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。

传送带上的第 i 个包裹的重量为 weights[i]。每一天，
我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。
*/

func shipWithinDays(weights []int, days int) int {
	if len(weights) == 0 {
		return 0
	}
	// 最快在一天内运完，最慢len(weights)天，一天内运完需要的最小载重量是所有weights之和，
	// len(weights)天运完，需要的最小载重量是weights中最大值。
	minWeight := 0
	maxWeight := 0
	for _, weight := range weights {
		if weight > minWeight {
			minWeight = weight
		}
		maxWeight += weight
	}
	for minWeight < maxWeight {
		mid := minWeight + ((maxWeight - minWeight) >> 1)
		d := getDays(weights, mid)
		if d == days {
			maxWeight = mid
		} else if d > days {
			minWeight = mid + 1
		} else if d < days {
			maxWeight = mid
		}
	}
	return minWeight
}

func getDays(weights []int, weight int) int {
	if len(weights) == 0 {
		return 0
	}
	day := 0
	w := weight
	for i := 0; i < len(weights); i++ {
		if w < weights[i] {
			day++
			w = weight
		}
		w -= weights[i]
	}
	return day + 1
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	fmt.Println(shipWithinDays(weights, days)) // 15
	weights = []int{1, 2, 3, 1, 1}
	days = 4
	fmt.Println(shipWithinDays(weights, days)) // 3
	weights = []int{3, 2, 2, 4, 1, 4}
	days = 3
	fmt.Println(shipWithinDays(weights, days)) // 6
}
