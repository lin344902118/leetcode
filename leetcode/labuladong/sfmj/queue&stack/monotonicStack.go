package main

import "fmt"

type MonotonicStack struct {
	data []int
}

func (s *MonotonicStack) Len() int {
	return len(s.data)
}

func (s *MonotonicStack) Push(val int) {
	if len(s.data) != 0 && s.data[len(s.data)-1] >= val {
		// 单调递增
		for i := len(s.data) - 2; i >= 0; i-- {
			if s.data[i] < val {
				s.data = s.data[:i+1]
				break
			}
		}
	}
	s.data = append(s.data, val)
}

func (s *MonotonicStack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func NewMonotonicStack() *MonotonicStack {
	return &MonotonicStack{}
}

func main() {
	s := NewMonotonicStack()
	s.Push(1)
	s.Push(3)
	s.Push(5)
	s.Push(4)
	s.Push(6)
	s.Push(8)
	s.Push(7)
	s.Push(7)
	fmt.Println(s.data)
}
