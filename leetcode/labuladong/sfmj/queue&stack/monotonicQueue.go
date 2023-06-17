package main

import "fmt"

type MonotonicQueue struct {
	datas []int
}

func (q *MonotonicQueue) max() int {
	if len(q.datas) > 0 {
		return q.datas[0]
	} else {
		return -1
	}
}

func (q *MonotonicQueue) push(num int) {
	for len(q.datas) > 0 && q.datas[len(q.datas)-1] < num {
		q.datas = q.datas[:len(q.datas)-1]
	}
	q.datas = append(q.datas, num)
}

func (q *MonotonicQueue) pop() int {
	if len(q.datas) > 0 {
		num := q.datas[0]
		q.datas = q.datas[1:]
		return num
	} else {
		return -1
	}
}

func NewMonitonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		datas: make([]int, 0),
	}
}

func main() {
	q := NewMonitonicQueue()
	q.push(1)
	q.push(3)
	q.push(-1)
	q.pop()
	fmt.Println(q.datas, q.max())
}
