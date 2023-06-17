package main

import "fmt"

// 单调栈实现，比普通实现块
type MaxQueue struct {
	datas []int
	queue []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		datas: make([]int, 0),
		queue: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) > 0 {
		return this.queue[0]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.datas = append(this.datas, value)
	for len(this.queue) > 0 && this.queue[len(this.queue)-1] < value {
		this.queue = this.queue[:len(this.queue)-1]
	}
	this.queue = append(this.queue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.datas) > 0 {
		num := this.datas[0]
		this.datas = this.datas[1:]
		if len(this.queue) > 0 && num == this.queue[0] {
			this.queue = this.queue[1:]
		}
		return num
	}
	return -1
}

// 普通实现，每次删除的时候需要额外维护max
// type MaxQueue struct {
// 	datas []int
// 	max   int
// }

// func Constructor() MaxQueue {
// 	return MaxQueue{
// 		datas: make([]int, 0),
// 		max:   -1,
// 	}
// }

// func (this *MaxQueue) Max_value() int {
// 	return this.max
// }

// func (this *MaxQueue) Push_back(value int) {
// 	if value > this.max {
// 		this.max = value
// 	}
// 	this.datas = append(this.datas, value)
// }

// func (this *MaxQueue) Pop_front() int {
// 	if len(this.datas) > 0 {
// 		num := this.datas[0]
// 		this.datas = this.datas[1:]
// 		if this.max == num {
// 			this.max = -1
// 			for i := 0; i < len(this.datas); i++ {
// 				if this.datas[i] > this.max {
// 					this.max = this.datas[i]
// 				}
// 			}
// 		}
// 		return num
// 	} else {
// 		return -1
// 	}
// }

func main() {
	q := Constructor()
	q.Push_back(1)
	q.Push_back(2)
	fmt.Println(q.Max_value())
	fmt.Println(q.Pop_front())
	fmt.Println(q.Max_value())
}
