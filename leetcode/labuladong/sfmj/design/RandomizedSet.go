package main

import (
	"fmt"
	"math/rand"
)

/*
380. O(1) 时间插入、删除和获取随机元素
实现RandomizedSet 类：
RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。
每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
*/

type RandomizedSet struct {
	cache map[int]int
	vals  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		cache: make(map[int]int, 0),
		vals:  make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exist := this.cache[val]
	if !exist {
		this.vals = append(this.vals, val)
		this.cache[val] = len(this.vals) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exist := this.cache[val]
	if exist {
		if index != len(this.vals)-1 {
			tmp := this.vals[len(this.vals)-1]
			this.vals[len(this.vals)-1] = this.vals[index]
			this.vals[index] = tmp
			this.cache[tmp] = index
		}
		this.vals = this.vals[:len(this.vals)-1]
		delete(this.cache, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(this.vals))
	return this.vals[randomIndex]
}

// type RandomizedSet struct {
// 	cache map[int]struct{}
// }

// func Constructor() RandomizedSet {
// 	return RandomizedSet{
// 		cache: make(map[int]struct{}, 0),
// 	}
// }

// func (this *RandomizedSet) Insert(val int) bool {
// 	_, exist := this.cache[val]
// 	if !exist {
// 		this.cache[val] = struct{}{}
// 		return true
// 	}
// 	return false
// }

// func (this *RandomizedSet) Remove(val int) bool {
// 	_, exist := this.cache[val]
// 	if exist {
// 		delete(this.cache, val)
// 		return true
// 	}
// 	return false
// }

// func (this *RandomizedSet) GetRandom() int {
// 	// 比下面快一点
// 	index := rand.Intn(len(this.cache))
// 	i := 0
// 	for key, _ := range this.cache {
// 		if i == index {
// 			return key
// 		}
// 		i++
// 	}
// 	return 0
// 	// keys := make([]int, 0, len(this.cache))
// 	// for k := range this.cache {
// 	// 	keys = append(keys, k)
// 	// }
// 	// index := rand.Intn(len(keys))
// 	// return keys[index]
// }

func main() {
	s := Constructor()
	s.Insert(0)
	fmt.Println(s.cache, s.vals)
	s.Insert(1)
	fmt.Println(s.cache, s.vals)
	s.Remove(0)
	fmt.Println(s.cache, s.vals)
	s.Insert(2)
	fmt.Println(s.cache, s.vals)
	s.Remove(1)
	fmt.Println(s.cache, s.vals)
	fmt.Println(s.GetRandom())
}
