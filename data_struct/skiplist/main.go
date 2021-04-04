package main

import (
	"fmt"
	"math/rand"
)

// http://redisbook.com/preview/skiplist/datastruct.html

// 跳表链表
type Skiplist struct {
	Head, Tail    *slNode
	Length, Level int
}

// 跳表层
type slLevel struct {
	// 前进指针
	Forward *slNode
	// 跨度: 当前层，该节点与前进指针指向的节点之间跨过了多少个底层数据节点
	Span int
}

// 跳表节点
type slNode struct {
	// 值对象
	Obj interface{}
	// 分值 用于排序用
	Score float64
	// 后退指针
	Backward *slNode
	// 当前节点的跳表层，如果在第i元素上有值，则说明该节点在第i层上有索引
	// 节点的第一层是level[0]，每个元素在此层都有值
	Levels []slLevel
}

func NewSkipList() *Skiplist {
	h := newSkipListNode(32, 0, nil)
	return &Skiplist{
		Head:   h,
		Tail:   nil,
		Length: 0,
		Level:  1,
	}
}
func newSkipListNode(level int, score float64, e interface{}) *slNode {
	h := &slNode{}
	h.Obj = e
	h.Score = score
	h.Levels = make([]slLevel, level)
	// 初始化结点层
	for _, l := range h.Levels {
		l.Forward = nil
		l.Span = 0
	}
	return h
}

/* 跳表插入的逻辑
 * 阶梯式向下遍历，对比score找到各层预插入的前置节点，
 * 同时记下每层跨越的span,向下递增 rank[i] = rank[i+1] + level[i].span
 * 初始化新节点
 * 随机成功一个层数level，0~level有索引，level~s.levels 无索引
 * 更新 0~level层的节点与span数，更新level~s.levels层的span
 * 更新回退节点与tail
 * s.length++
 */
func (s *Skiplist) Insert(e int, score float64) *slNode {
	// update记录每一层插入节点的前置节点
	update := make([]*slNode, 32)
	rank := make([]int, 32)
	x := s.Head
	// 寻找在各个层插入的位置
	// 阶梯式向下遍历
	for i := s.Level - 1; i >= 0; i-- {
		if i == s.Level-1 {
			rank[i] = 0
		} else {
			// 各个层的rank值一层层累积
			// 最后rank[0]的值加一就是新节点的前置节点的排位
			rank[i] = rank[i+1]
		}
		// 遍历当前层，找到小于score的节点，记录跨越节点数，
		// 标记当前节点，用于插入新节点，然后进入下一层
		for x.Levels[i].Forward != nil &&
			(x.Levels[i].Forward.Score < score) {
			// 记录沿途跨越了多少个节点
			rank[i] += x.Levels[i].Span
			// 移动下一指针
			x = x.Levels[i].Forward
		}
		// 记录将要和新节点相连的节点
		update[i] = x
	}
	level := randReturnLevel()
	if level > s.Level {
		for i := s.Level; i < level; i++ {
			rank[i] = 0
			update[i] = s.Head
			update[i].Levels[i].Span = s.Length
		}
		s.Level = level
	}
	x = newSkipListNode(level, score, e)
	// 插入新节点 并更新 0~level层的span
	for i := 0; i < level; i++ {
		x.Levels[i].Forward = update[i].Levels[i].Forward
		update[i].Levels[i].Forward = x
		// 更新前置节点update[i]对应的span = 底层数组节点位置 - 第i层索引节点位置 + 1
		update[i].Levels[i].Span = rank[0] - rank[i] + 1
	}
	// 更新没有索引层的span值
	// 这些节点直接从表头指向新节点
	for i := level; i < s.Level; i++ {
		update[i].Levels[i].Span++
	}

	// 更新新节点回退指针
	if update[0] == s.Head {
		// 第一个元素，回退指针指向空
		x.Backward = nil
	} else {
		// 回退指针指向底层数组的上一个元素
		x.Backward = update[0]
	}
	// 如果x的下个节点不是空，则更新其回退节点为x
	if x.Levels[0].Forward != nil {
		x.Levels[0].Forward.Backward = x
	} else {
		// 如果x不具备下个节点，说明他是最后一个节点
		s.Tail = x
	}
	s.Length++
	return x
}

/* 跳表删除逻辑
 * 阶梯式向下遍历，标记需要删除的节点的前置节点
 * 遍历第0层，找到需要更新的节点
 * 从0层向上遍历，逐层删除节点，
 * 索引层有节点时删除节点相当于把节点两边的span串在一起，没有则span--
 * 底层双向链表进行删除
 * 判断是否因为删除节点而出现空层
 */
func (s *Skiplist) Delete(e int, score float64) {
	update := make([]*slNode, 32)
	x := s.Head
	// 阶梯式向下遍历，标记需要删除的节点的前置节点
	for i := s.Level; i >= 0; i-- {
		if x.Levels[i].Forward != nil && (x.Levels[i].Forward.Score < score) {
			x = x.Levels[i].Forward
		}
		update[i] = x
	}
	// 遍历结束后 x是需要删除节点的前置节点
	// 遍历第0层，找到需要更新的节点,
	x = x.Levels[0].Forward
	if x != nil && score == x.Score {
		// 更新各个层
		for i := 0; i < s.Level; i++ {
			if update[i].Levels[i].Forward == x {
				// 删除掉x节点相当于将x节点两边的span串在一起
				update[i].Levels[i].Span += x.Levels[i].Span - 1
				update[i].Levels[i].Forward = x.Levels[i].Forward
			} else {
				// 第i层没有x节点，直接span--
				update[i].Levels[i].Span--
			}
		}
		// 底层双向链表进行删除
		if x.Levels[0].Forward != nil {
			x.Levels[0].Forward.Backward = x.Backward
		} else {
			// 尾节点
			s.Tail = x.Backward
		}
		// 判断是否因为删除节点而出现空层
		for s.Level > 1 && s.Head.Levels[s.Level-1].Forward == nil {
			s.Level--
		}
	}
	s.Length--
}
func main() {
	sl := NewSkipList()
	sl.Insert(2, 10)
	sl.Insert(1, 20)
	sl.Insert(3, 30)
	sl.Insert(6, 40)
	fmt.Printf("level %v    length %v \n", sl.Level, sl.Length)
	x := sl.Head.Levels[0].Forward
	for x != nil {
		fmt.Println(x.Obj.(int))
		x = x.Levels[0].Forward
	}

	sl.Delete(1, 20)
	sl.Delete(3, 30)
	x = sl.Head.Levels[0].Forward
	for x != nil {
		fmt.Println(x.Obj.(int))
		x = x.Levels[0].Forward
	}
}

// 返回一个随机值用作初始化跳跃节点的层数
func randReturnLevel() int {
	//rand.Seed(2)
	return rand.Intn(33)
}
