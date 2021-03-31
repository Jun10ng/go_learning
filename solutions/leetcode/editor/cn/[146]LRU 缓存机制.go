package main
//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。 
//
// 
// 
// 实现 LRUCache 类： 
//
// 
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存 
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。 
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上
//限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。 
// 
//
// 
// 
// 
//
// 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？ 
//
// 
//
// 示例： 
//
// 
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
// 
//
// 
//
// 提示： 
//
// 
// 1 <= capacity <= 3000 
// 0 <= key <= 3000 
// 0 <= value <= 104 
// 最多调用 3 * 104 次 get 和 put 
// 
// Related Topics 设计 
// 👍 1291 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

type LRUCache struct {
	cap int
	m map[int]*node
	head *node
	tail *node
}
type node struct {
	key int
	value int
	next *node
	pre *node
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		make(map[int]*node,capacity),
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
	if this.tail.value == n.value{
		this.tail = this.tail.pre
	}
	pre := n.pre
	nxt := n.next
	pre.next = nxt
	if nxt!=nil {
		nxt.pre = pre
	}
	n.next = this.head
	this.head.pre = n
	this.head = n
	return n.value
}

func (this *LRUCache) Put(key int, value int)  {
	n := &node{
		key: key,
		value: value,
	}
	if len(this.m)==0{
		this.tail = n
	}else {
		n.next = this.head
		this.head.pre = n
	}
	this.m[key] = n
	this.head = n
	if  len(this.m) > this.cap  {
		d := this.tail
		this.tail = this.tail.pre
		if d != nil {
			delete(this.m,d.key)
			d = nil
		}
	}
}





/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
