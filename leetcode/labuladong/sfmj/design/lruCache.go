package main

import (
	"fmt"
)

/*
146. LRU 缓存
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，
则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

type entry struct {
	key int
	val int
}

type Node struct {
	prev *Node
	next *Node
	val  entry
}

type DoubleLinkedList struct {
	head    *Node
	tail    *Node
	curSize int
}

// 创建新的双链表
func NewDoubleLinkedList() *DoubleLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoubleLinkedList{
		head:    head,
		tail:    tail,
		curSize: 0,
	}
}

// 往双链表尾部添加元素
func (d *DoubleLinkedList) push(ele *Node) {
	ele.prev = d.tail.prev
	ele.next = d.tail
	d.tail.prev.next = ele
	d.tail.prev = ele
	d.curSize++
}

// 往双链表头部删除元素
func (d *DoubleLinkedList) pop() *Node {
	if d.curSize == 0 {
		return &Node{}
	}
	ele := d.head.next
	ele.next.prev = d.head
	d.head.next = ele.next
	d.curSize--
	return ele
}

// 删除双链表指定元素
func (d *DoubleLinkedList) popElement(ele *Node) {
	if d.curSize == 0 {
		return
	}
	tmp := d.head.next
	for tmp != d.tail {
		if tmp.val == ele.val {
			tmp.prev.next = tmp.next
			tmp.next.prev = tmp.prev
			d.curSize--
			break
		}
		tmp = tmp.next
	}
}

type LRUCache struct {
	ll    *DoubleLinkedList
	cache map[int]*Node
	cap   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		ll:    NewDoubleLinkedList(),
		cache: make(map[int]*Node, capacity),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	ele, exist := this.cache[key]
	if exist {
		// 将元素移动到队尾
		this.ll.popElement(ele)
		this.ll.push(ele)
		return ele.val.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	ele, exist := this.cache[key]
	if exist {
		this.ll.popElement(ele)
	}
	if this.ll.curSize == this.cap {
		// 删除最近最少使用的
		e := this.ll.pop()
		delete(this.cache, e.val.key)
	}
	newNode := &Node{
		val: entry{
			key: key,
			val: value,
		},
	}
	this.cache[key] = newNode
	this.ll.push(newNode)
}

// type LRUCache struct {
// 	ll    *list.List
// 	cache map[int]*list.Element
// 	cap   int
// }

// type data struct {
// 	key int
// 	val int
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{
// 		ll:    list.New(),
// 		cache: make(map[int]*list.Element, capacity),
// 		cap:   capacity,
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	ele, exist := this.cache[key]
// 	if exist {
// 		// 将元素移动到队尾
// 		this.ll.MoveToFront(ele)
// 		e := ele.Value.(data)
// 		return e.val
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	ele, exist := this.cache[key]
// 	if exist {
// 		this.ll.Remove(ele)
// 	}
// 	if this.ll.Len() == this.cap {
// 		// 删除最近最少使用的
// 		e := this.ll.Back()
// 		d := e.Value.(data)
// 		delete(this.cache, d.key)
// 		this.ll.Remove(e)
// 	}
// 	newE := data{
// 		key: key,
// 		val: value,
// 	}
// 	e := this.ll.PushFront(newE)
// 	this.cache[key] = e
// }

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))

	lru = Constructor(2)
	fmt.Println(lru.Get(2))
	lru.Put(2, 6)
	fmt.Println(lru.Get(1))
	lru.Put(1, 5)
	lru.Put(1, 2)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
}
