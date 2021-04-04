package main

import "fmt"

func main() {

	lRUCache := Constructor(1)
	lRUCache.Put(2, 2)           // 缓存是 {1=1}
	lRUCache.Put(1, 1)           // 缓存是 {1=1, 2=2}
	fmt.Println(lRUCache.Get(2)) // 返回 1
	//lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	//fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
	//lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	//fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
	//fmt.Println(lRUCache.Get(3)) // 返回 3
	//fmt.Println(lRUCache.Get(4)) // 返回 4

}

type LRUCache struct {
	cap  int
	m    map[int]*node
	head *node
	tail *node
}
type node struct {
	key   int
	value int
	next  *node
	pre   *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		make(map[int]*node, capacity),
		&node{},
		&node{},
	}
}

func (this *LRUCache) Get(key int) int {
	var n *node
	var ok bool
	if n, ok = this.m[key]; !ok {
		return -1
	}
	// 提前
	if this.tail.value == n.value {
		this.tail = this.tail.pre
	}
	pre := n.pre
	nxt := n.next
	pre.next = nxt
	if nxt != nil {
		nxt.pre = pre
	}
	n.next = this.head
	this.head.pre = n
	this.head = n
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	n := &node{
		key:   key,
		value: value,
	}
	if len(this.m) == 0 {
		this.tail = n
	} else {
		n.next = this.head
		this.head.pre = n
	}
	this.m[key] = n
	this.head = n
	if len(this.m) > this.cap {
		d := this.tail
		this.tail = this.tail.pre
		if d != nil {
			delete(this.m, d.key)
			d = nil
		}
	}
}
