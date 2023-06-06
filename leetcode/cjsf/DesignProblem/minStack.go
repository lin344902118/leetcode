package main

import (
	"fmt"
	"math"
)

/*
  最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。
*/

type MinStack struct {
	data []int
	min  int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  math.MaxInt32,
	}
}

func (this *MinStack) Push(val int) {
	if val < this.min {
		this.min = val
	}
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	if len(this.data) == 1 {
		this.min = math.MaxInt32
		this.data = []int{}
		return
	}
	if this.min == this.data[len(this.data)-1] {
		min := this.data[0]
		for i := 1; i < len(this.data)-1; i++ {
			if this.data[i] < min {
				min = this.data[i]
			}
		}
		this.min = min
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	fmt.Println(m.data)
	fmt.Println(m.Top())
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.data)
	fmt.Println(m.GetMin())
	fmt.Println(m.Top())
}
