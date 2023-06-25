package main

import "fmt"

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

func printLinkedlist(ll *DoubleLinkedList) {
	tmp := ll.head.next
	for tmp != ll.tail {
		fmt.Println(tmp.val)
		tmp = tmp.next
	}
}

func main() {
	ll := NewDoubleLinkedList()
	ll.push(&Node{val: entry{key: 1, val: 1}})
	ll.push(&Node{val: entry{key: 2, val: 2}})
	fmt.Println(ll.curSize)
	printLinkedlist(ll)
	ll.popElement(&Node{val: entry{key: 3, val: 3}})
	fmt.Println(ll.curSize)
	printLinkedlist(ll)
	ll.push(&Node{val: entry{key: 3, val: 3}})
	ll.push(&Node{val: entry{key: 4, val: 4}})
	fmt.Println(ll.curSize)
	printLinkedlist(ll)
	ll.pop()
	fmt.Println(ll.curSize)
	printLinkedlist(ll)
	ll.popElement(&Node{val: entry{key: 3, val: 3}})
	fmt.Println(ll.curSize)
	printLinkedlist(ll)
}
