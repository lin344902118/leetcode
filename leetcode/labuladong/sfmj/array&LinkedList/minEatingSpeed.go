package main

import "fmt"

/*
875. 爱吃香蕉的珂珂
珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。
珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。
如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。
*/

func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 0 {
		return 0
	}
	maxSpeed := 0
	for i := 0; i < len(piles); i++ {
		maxSpeed += piles[i]
	}
	minSpeed := maxSpeed / h
	if minSpeed < 1 {
		minSpeed = 1
	}
	for minSpeed < maxSpeed {
		mid := minSpeed + ((maxSpeed - minSpeed) >> 1)
		hour := getHour(piles, mid)
		if hour == h {
			maxSpeed = mid
		} else if hour < h {
			maxSpeed = mid
		} else if hour > h {
			minSpeed = mid + 1
		}
	}
	return minSpeed
}

// 给一堆香蕉和每小时吃的速度，求需要多少小时
func getHour(piles []int, speed int) (hour int) {
	if len(piles) == 0 {
		return hour
	}
	for _, pile := range piles {
		if pile < speed {
			hour++
			continue
		}
		hour += pile / speed
		if pile%speed != 0 {
			hour++
		}
	}
	return hour
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	// fmt.Println(minEatingSpeed(piles, h)) // 4
	// piles = []int{30, 11, 23, 5, 20}
	// h = 5
	// fmt.Println(minEatingSpeed(piles, h)) // 30
	// h = 6
	// fmt.Println(minEatingSpeed(piles, h)) // 23
	piles = []int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}
	h = 823855818
	fmt.Println(minEatingSpeed(piles, h))
}
