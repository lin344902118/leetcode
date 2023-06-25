package main

import (
	"fmt"
	"math"
)

/*
460. LFU 缓存
请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
int Get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
void Put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。
当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。
在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 Put 操作)。
对缓存中的键执行 Get 或 Put 操作，使用计数器的值将会递增。

函数 Get 和 Put 必须以 O(1) 的平均时间复杂度运行。
*/

type entry struct {
	key   int
	val   int
	count int
}

type Node struct {
	prev *Node
	next *Node
	val  *entry
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
	ele.prev.next = ele.next
	ele.next.prev = ele.prev
	d.curSize--
	return
}

// 将指定元素移动到队尾
func (d *DoubleLinkedList) moveToBack(ele *Node) {
	if d.curSize == 0 {
		return
	}
	ele.prev.next = ele.next
	ele.next.prev = ele.prev
	d.tail.prev.next = ele
	ele.prev = d.tail.prev
	ele.next = d.tail
	d.tail.prev = ele
}

// 将指定元素移动到队头
func (d *DoubleLinkedList) moveToFont(ele *Node) {
	if d.curSize == 0 {
		return
	}
	ele.prev.next = ele.next
	ele.next.prev = ele.prev
	ele.next = d.head.next
	d.head.next.prev = ele
	ele.prev = d.head
	d.head.next = ele
}

// 获取双链表队头元素
func (d *DoubleLinkedList) GetFont() *Node {
	return d.head.next
}

// 获取双链表队尾元素
func (d *DoubleLinkedList) GetBack() *Node {
	return d.tail.prev
}

func (d *DoubleLinkedList) Len() int {
	return d.curSize
}

type LFUCache struct {
	// 缓存数据 key是传入key，value是node
	cache map[int]*Node
	// 缓存最大容量
	cap int
	// 缓存访问次数，key是访问次数，value是元素链表
	freq map[int]*DoubleLinkedList
	// 当前存在元素个数
	curSize int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:   make(map[int]*Node, capacity),
		cap:     capacity,
		freq:    make(map[int]*DoubleLinkedList, 0),
		curSize: 0,
	}
}

// 判断key是否在缓存中，如果不是，返回-1.如果是，获取对应的节点。根据节点的访问次数，将其从访问map链表中删除。增加访问次数，并添加到访问map链表中
func (this *LFUCache) Get(key int) int {
	node, exist := this.cache[key]
	if exist {
		e := node.val
		oldList := this.freq[e.count]
		oldList.popElement(node)
		e.count++
		_, exist := this.freq[e.count]
		if !exist {
			this.freq[e.count] = NewDoubleLinkedList()
		}
		this.freq[e.count].push(node)
		return e.val
	}
	return -1
}

// 如果存在将key从旧链表删除，并且从访问map链表中删除。如果容量不够，找到访问次数最少的，从访问map中删除，并从缓存中删除。
// 将新值加入缓存，访问次数加一，并添加到访问map链表中
func (this *LFUCache) Put(key int, value int) {
	// 继承旧key的count
	var count int
	node, exist := this.cache[key]
	if exist {
		count = node.val.count
		oldList := this.freq[count]
		oldList.popElement(node)
	}
	if this.curSize == this.cap && !exist {
		minCount := math.MaxInt32
		for key, l := range this.freq {
			if key < minCount && l.Len() > 0 {
				minCount = key
			}
		}
		list := this.freq[minCount]
		deleteNode := list.GetFont()
		list.pop()
		e := deleteNode.val
		delete(this.cache, e.key)
		this.curSize--
	}
	if !exist {
		this.curSize++
	}
	count++
	n := &Node{
		val: &entry{
			key:   key,
			val:   value,
			count: count,
		},
	}
	_, e := this.freq[count]
	if !e {
		this.freq[count] = NewDoubleLinkedList()
	}
	this.freq[count].push(n)
	this.cache[key] = n
}

func printLfuCache(lfu *LFUCache) {
	fmt.Printf("cache:%+v, size:%d, cap:%d\n", lfu.cache, lfu.curSize, lfu.cap)
	fmt.Println("freqMap")
	for key, value := range lfu.freq {
		fmt.Println("count", key)
		tmp := value.head.next
		for i := 0; i < value.Len(); i++ {
			fmt.Printf("%d", tmp.val)
			fmt.Printf(" ")
			tmp = tmp.next
		}
		fmt.Printf("\n")
	}
}

func main() {
	// lfu := Constructor(2)
	// lfu.Put(1, 1)           // [1:1]
	// lfu.Put(2, 2)           // [1:1 2:1]
	// fmt.Println(lfu.Get(1)) // 1  [2:1 1:2]
	// lfu.Put(3, 3)           // [1:2 3:1]
	// fmt.Println(lfu.Get(2)) // -1 [1:2 3:1]
	// fmt.Println(lfu.Get(3)) // 3 [1:2 3:2]
	// lfu.Put(4, 4)           // [3:2 4:1]
	// fmt.Println(lfu.Get(1)) // -1 [3:2 4:1]
	// fmt.Println(lfu.Get(3)) // 3 [4:1 3:3]
	// fmt.Println(lfu.Get(4)) // 4 [3:3 4:2]
	// lfu := Constructor(2)
	// fmt.Println(lfu.Get(2)) // -1
	// lfu.Put(2, 6)           // [2:1]
	// fmt.Println(lfu.Get(1)) // -1
	// lfu.Put(1, 5)           // [2:1 1:1]
	// lfu.Put(1, 2)           // [2:1 1:2]
	// fmt.Println(lfu.Get(1)) // 2
	// fmt.Println(lfu.Get(2)) // 6
	lfu := Constructor(10)
	lfu.Put(10, 13)
	lfu.Put(3, 17)
	lfu.Put(6, 11)
	lfu.Put(10, 5)
	lfu.Put(9, 10)
	fmt.Println(lfu.Get(13))
	lfu.Put(2, 19)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
	lfu.Put(5, 25)
	fmt.Println(lfu.Get(8))
	lfu.Put(9, 22)
	lfu.Put(5, 5)
	lfu.Put(1, 30)
	fmt.Println(lfu.Get(11))
	lfu.Put(9, 12)
	fmt.Println(lfu.Get(7))
	fmt.Println(lfu.Get(5))
	fmt.Println(lfu.Get(8))
	fmt.Println(lfu.Get(9))
	lfu.Put(4, 30)
	lfu.Put(9, 3)
	fmt.Println(lfu.Get(9))
	fmt.Println(lfu.Get(10))
	fmt.Println(lfu.Get(10))
	lfu.Put(6, 14)
	lfu.Put(3, 1)
	fmt.Println(lfu.Get(3))
	lfu.Put(10, 11)
	fmt.Println(lfu.Get(8))
	lfu.Put(2, 14)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(5))
	fmt.Println(lfu.Get(4))
	lfu.Put(11, 4)
	lfu.Put(12, 24)
	lfu.Put(5, 18)
	fmt.Println(lfu.Get(13))
	lfu.Put(7, 23)
	fmt.Println(lfu.Get(8))
	fmt.Println(lfu.Get(12))

	lfu.Put(3, 27)
	lfu.Put(2, 12)
	fmt.Println(lfu.Get(5))
	lfu.Put(2, 9)
	lfu.Put(13, 4)
	lfu.Put(8, 18)
	lfu.Put(1, 7)
	fmt.Println(lfu.Get(6))

	lfu.Put(9, 29)
	lfu.Put(8, 21)
	fmt.Println(lfu.Get(5))
	lfu.Put(6, 30)
	lfu.Put(1, 12)
	fmt.Println(lfu.Get(10))
	lfu.Put(4, 15)
	// printLfuCache(&lfu)
	lfu.Put(7, 22)
	lfu.Put(11, 26)
	lfu.Put(8, 17)
	lfu.Put(9, 29)
	fmt.Println(lfu.Get(5))
	lfu.Put(3, 4)
	lfu.Put(11, 30)
	fmt.Println(lfu.Get(12))
	lfu.Put(4, 29)
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(9))
	fmt.Println(lfu.Get(6))
	lfu.Put(3, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(10))
	lfu.Put(3, 29)
	lfu.Put(10, 28)
	lfu.Put(1, 20)
	lfu.Put(11, 13)
	fmt.Println(lfu.Get(3))
	lfu.Put(3, 12)
	lfu.Put(3, 8)
	lfu.Put(10, 9)
	lfu.Put(3, 26)
	fmt.Println(lfu.Get(8))
	fmt.Println(lfu.Get(7))
	fmt.Println(lfu.Get(5))
	lfu.Put(13, 17)
	lfu.Put(2, 27)
	lfu.Put(11, 15)
	fmt.Println(lfu.Get(12))
	lfu.Put(9, 19)
	lfu.Put(2, 15)
	lfu.Put(3, 16)
	fmt.Println(lfu.Get(1))
	lfu.Put(12, 17)
	lfu.Put(9, 1)
	lfu.Put(6, 19)
	fmt.Println(lfu.Get(4))
	fmt.Println(lfu.Get(5))
	fmt.Println(lfu.Get(5))
	lfu.Put(8, 1)
	lfu.Put(11, 7)
	lfu.Put(5, 2)
	lfu.Put(9, 28)
	fmt.Println(lfu.Get(1))
	lfu.Put(2, 2)
	lfu.Put(7, 4)
	lfu.Put(4, 22)
	lfu.Put(7, 24)
	lfu.Put(9, 26)
	lfu.Put(13, 28)
	lfu.Put(11, 26)
	printLfuCache(&lfu)
}
