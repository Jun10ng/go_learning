package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	pre := []int{1, 2}
	in := []int{2, 1}
	root := buildTree(pre, in)
	fmt.Println(root)

}

// 前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	build := func(preorder []int, inorder []int) *TreeNode {
		return nil
	}
	build = func(pre []int, in []int) *TreeNode {
		if len(pre) == 0 {
			return nil
		}
		r := &TreeNode{
			Val: pre[0],
		}
		if len(in) == 1 {
			return r
		}
		idx := findIdx(in, pre[0]) // 左子树+根节点 节点数
		r.Left = build(pre[1:idx+1], in[:idx])
		r.Right = build(pre[idx+1:], in[idx+1:])
		return r
	}
	return build(preorder, inorder)
}
func findIdx(s []int, e int) int {
	for i, ie := range s {
		if ie == e {
			return i
		}
	}
	return -1
}

type Codec struct {
	vs []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ds := this.doSerialize(root)
	return strings.Join(ds, ",")
}

func (this *Codec) doSerialize(root *TreeNode) []string {
	if root == nil {
		return []string{""}
	}
	ret := []string{strconv.Itoa(root.Val)}
	ret = append(ret, this.serialize(root.Left))
	ret = append(ret, this.serialize(root.Right))
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.vs = strings.Split(data, ",")
	return this.doDeserialize()
}
func (this *Codec) doDeserialize() *TreeNode {
	if this.vs[0] == "" {
		this.vs = this.vs[1:]
		return nil
	}

	v, _ := strconv.Atoi(this.vs[0])
	this.vs = this.vs[1:]
	n := &TreeNode{
		Val:   v,
		Left:  this.doDeserialize(),
		Right: this.doDeserialize(),
	}
	return n
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = root.Left, root.Right
	root.Right = invertTree(root.Right)
	root.Left = invertTree(root.Left)
	return root
}

// 1 3
// 0 2 1 3 4 5
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	// new head
	nh := &ListNode{0, head}
	pre := nh
	for count := 0; pre.Next != nil && count < left-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return nh.Next
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	ret = append(ret, inorderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)
	return ret
}
func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	m := make(map[string]int)
	for _, v := range strs {
		vs := strings.Split(v, "")
		sort.Strings(vs)
		key := strings.Join(vs, "")
		index, ok := 0, false
		if index, ok = m[key]; ok {
		} else {
			ret = append(ret, []string{})
			m[key] = len(m)
			index = m[key]
		}
		ret[index] = append(ret[index], v)
	}
	return ret
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = nums2
	}
	i, j, k := m-1, n-1, n+m-1
	for k >= 0 {
		if j > 0 || i < 0 {
			if j >= 0 && i < 0 {
				for l := 0; l <= j; l++ {
					nums1[l] = nums2[l]
				}
			}
			fmt.Println(nums1)
			return
		}
		if nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			k--
			j--
			continue
		}
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
			continue
		}
	}
	fmt.Println(nums1)
}

////输入：l1 = [1,2,4], l2 = [3]
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists(l1.Next, l2)
//		return l1
//	}
//	l2.Next = mergeTwoLists(l1, l2.Next)
//	return l2
//}
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 反转法，时间复杂度和空间复杂度都是最优的
func rotate(nums []int, k int) {
	n := len(nums)
	if k == n {
		return
	}

	// 翻转函数
	reserve := func(ns []int) {
		for l, r := 0, len(ns)-1; l < r; {
			ns[l], ns[r] = ns[r], ns[l]
			l++
			r--
		}
	}
	reserve(nums)
	reserve(nums[:k])
	reserve(nums[k:])
}
func rotate1(nums []int, k int) {
	i := 0
	n := len(nums)
	for i < k {
		t := nums[n-1]
		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = t
		i++
	}
}

// 双指针法
// 以q为游标遍历数组，把不同的元素丢到p指针的左侧
func removeDuplicates(nums []int) int {
	p, q := 0, 1
	for q < len(nums) {
		if nums[p] == nums[q] {
			q++
		} else {
			nums[p+1] = nums[q]
			p++
			q++
		}
	}
	return p + 1
}
func removeDuplicates1(nums []int) int {
	/*
		//输入：nums = [0,0,1,1,1,2,2,3,3,4]
		//输出：5, nums = [0,1,2,3,4]
	*/
	for i := 0; i < len(nums); {
		v := nums[i]
		if i > 0 && v == nums[i-1] {
			// remove
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
func moveZeroes(nums []int) {
	length := len(nums)
	for i := 0; i < length; {
		if nums[i] != 0 {
			i++
			continue
		}
		nums = append(nums[:i], nums[i+1:]...)
		nums = append(nums, 0)
		length--
	}
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	//m := make(map[int]int,len(nums))
	//for index:=0;index<len(nums);index++{
	//	value := nums[index]
	//	m[value] = index
	//}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ret
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if s > 0 {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}

	}
	return ret
}
func maxArea(height []int) int {
	// 双指针
	l, r := 0, len(height)-1
	area := 0
	for l < r {
		a := min2(height[l], height[r]) * (r - l)
		area = max2(a, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
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

//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		capacity,
//		make(map[int]*node, capacity),
//		&node{},
//		&node{},
//	}
//}

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

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = 9999999
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min2(dp[i], dp[i-coin]+1)
		}
	}
	return dp[amount]
}

// 输出num的排列组合
func helper(str string) []string {
	strs := strings.Split(str, "")
	sort.Strings(strs)
	str = strings.Join(strs, "")
	strList := make([]string, 0)

	vis := make([]bool, len(str))
	trace := func(cur string) {}
	trace = func(cur string) {
		if len(cur) == len(str) {
			strList = append(strList, cur)
			return
		}
		for i := 0; i < len(str); i++ {
			se := string(str[i])
			if !vis[i] && !(i != 0 && vis[i-1] && str[i] == str[i-1]) {
				vis[i] = true
				trace(cur + se)
				vis[i] = false
			}
		}
	}
	trace("")
	return strList
}

type IntHeap []int

func (h *IntHeap) Len() int {
	return len((*h))
}
func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *IntHeap) Push(e interface{}) {
	*h = append(*h, e.(int))
}
func (h *IntHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return ret
}
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i, e := range nums {
		heap.Push(h, e)
		if i >= k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	findKth := func(k int) float64 {
		i, j := 0, 0
		for {
			//某个序列为空
			if i == len(nums1) {
				return float64(nums2[j+k-1])
			}
			if j == len(nums2) {
				return float64(nums1[i+k-1])
			}
			if k == 1 {
				return float64(min2(nums1[i], nums2[j]))
			}
			halfK := k/2 - 1
			newi := min2(i+halfK, len(nums1)-1)
			newj := min2(j+halfK, len(nums2)-1)
			if nums1[newi] < nums2[newj] {
				k = k - (newi - i + 1)
				i = newi + 1
			} else {
				k = k - (newj - j + 1)
				j = newj + 1
			}
		}
	}
	l1, l2 := len(nums1), len(nums2)
	isEven := (l1+l2)&1 == 0
	mid := (l1 + l2) / 2
	if isEven {
		return (findKth(mid+1) + findKth(mid)) / 2
	}
	return findKth(mid + 1)
}

func longestConsecutive(nums []int) int {
	u := NewUnionFindSet(nums)
	pre := make(map[int]int, len(nums))
	for i, e := range nums {
		if _, exist := pre[e]; exist {
			continue
		}
		if prei, ok := pre[e+1]; ok {
			u.Union(i, prei)
		}
		if prei, ok := pre[e-1]; ok {
			u.Union(i, prei)
		}
		pre[e] = i
	}
	return u.LC
}

type UnionFindSet struct {
	M      int
	Nums   []int
	Parent []int
	Rank   []int
	LC     int
}

func NewUnionFindSet(nums []int) *UnionFindSet {
	m := len(nums)
	if m == 0 {
		return &UnionFindSet{}
	}

	p := make([]int, m)
	r := make([]int, m)
	for i, _ := range p {
		p[i] = i //每个元素的父亲都是自己
		r[i] = 1
	}
	return &UnionFindSet{
		M:      m,
		Nums:   nums,
		Parent: p,
		Rank:   r,
		LC:     1,
	}
}
func (u *UnionFindSet) find(x int) int {
	px := u.Parent[x]
	if u.Parent[x] == x { //找到父节点
		return x
	}
	return u.find(px)
}
func (u *UnionFindSet) Union(x, y int) {
	px := u.find(x)
	py := u.find(y)
	if px != py {
		if u.Nums[x] > u.Nums[y] {
			u.Parent[x] = py
			u.Rank[py] += u.Rank[x]
			u.LC = max2(u.Rank[py], u.LC)
		} else {
			u.Parent[y] = px
			u.Rank[px] += u.Rank[y]
			u.LC = max2(u.Rank[px], u.LC)
		}
	}
}

//func findCircleNum(isConnected [][]int) int {
//	u := NewUnionFindSet(isConnected)
//	for i := 0; i < u.M; i++ {
//		for j := i + 1; j < u.M; j++ {
//			if isConnected[i][j] == 1 {
//				u.Union(i, j)
//			}
//		}
//	}
//	return u.Cnt
//}
//
//type UnionFindSet struct {
//	M, N, Cnt int
//	Grid      [][]int
//	Parent    []int
//	Rank      []int //深度
//}
//
//func NewUnionFindSet(g [][]int) *UnionFindSet {
//	m := len(g)
//	n := len(g[0])
//	cnt := 0
//	if n == 0 {
//		return &UnionFindSet{}
//	}
//	p := make([]int, m)
//	r := make([]int, m)
//	for i, _ := range p {
//		p[i] = i
//		cnt++
//	}
//
//	return &UnionFindSet{
//		Grid:   g,
//		Parent: p,
//		Rank:   r,
//		M:      m,
//		N:      n,
//		Cnt:    cnt,
//	}
//}
//func (u *UnionFindSet) Find(index int) int {
//	if u.Parent[index] == index { //找到根节点
//		return index
//	}
//	return u.Find(u.Parent[index]) // 递归子节点
//}
//
//func (u *UnionFindSet) Union(x, y int) {
//	rootx := u.Find(x)
//	rooty := u.Find(y)
//	if rootx != rooty {
//		//两个节点 根节点不一样，
//		//但是两个节点相邻，所以进行union,cnt--
//		//把深度(rank)小的一方指向深度大的
//		if u.Rank[rootx] > u.Rank[rooty] {
//			u.Parent[rooty] = rootx
//		} else if u.Rank[rooty] > u.Rank[rootx] {
//			u.Parent[rootx] = rooty
//		} else {
//			// 两个节点深度相同，选择让rooty指向rootx，rootx深度+1
//			u.Parent[rooty] = rootx
//			u.Rank[rootx]++
//		}
//		u.Cnt--
//	}
//}

//
//func numIslands(grid [][]byte) int {
//	u := NewUnionFindSet(grid)
//	for i, ei := range grid {
//		for j, ej := range ei {
//			if ej == '1' {
//				for _, v := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
//					vi, vj := i+v[0], j+v[1]
//					if vi >= 0 && vj >= 0 && vi < u.M && vj < u.N && grid[vi][vj] == '1' {
//						u.Union(i*u.N+j, vi*u.N+vj)
//					}
//				}
//			}
//		}
//	}
//	return u.Cnt
//}
//
//type UnionFindSet struct {
//	M, N, Cnt int
//	Grid      [][]byte
//	Parent    []int
//	Rank      []int //深度
//}
//
//func NewUnionFindSet(g [][]byte) *UnionFindSet {
//	m := len(g)
//	n := len(g[0])
//	cnt := 0
//	if n == 0 {
//		return &UnionFindSet{}
//	}
//	p := make([]int, m*n)
//	r := make([]int, m*n)
//	for i, _ := range p {
//		p[i] = -1
//	}
//	for i, ei := range g {
//		for j, ej := range ei {
//			if ej == '1' {
//				p[i*n+j] = i*n + j
//				cnt++
//			}
//		}
//	}
//	return &UnionFindSet{
//		Grid:   g,
//		Parent: p,
//		Rank:   r,
//		M:      m,
//		N:      n,
//		Cnt:    cnt,
//	}
//}
//func (u *UnionFindSet) Find(index int) int {
//	if u.Parent[index] == index { //找到根节点
//		return index
//	}
//	return u.Find(u.Parent[index]) // 递归子节点
//}
//
//func (u *UnionFindSet) Union(x, y int) {
//	rootx := u.Find(x)
//	rooty := u.Find(y)
//	if rootx != rooty {
//		//两个节点 根节点不一样，
//		//但是两个节点相邻，所以进行union,cnt--
//		//把深度(rank)小的一方指向深度大的
//		if u.Rank[rootx] > u.Rank[rooty] {
//			u.Parent[rooty] = rootx
//		} else if u.Rank[rooty] > u.Rank[rootx] {
//			u.Parent[rootx] = rooty
//		} else {
//			// 两个节点深度相同，选择让rooty指向rootx，rootx深度+1
//			u.Parent[rooty] = rootx
//			u.Rank[rootx]++
//		}
//		u.Cnt--
//	}
//}

type Mem [][]int

func (m Mem) Repalce(i, j int) int {
	return m[i-1][j-1]
}
func (m Mem) Insert(i, j int) int {
	return m[i][j-1] // ro  ros
}
func (m Mem) Delete(i, j int) int {
	return m[i-1][j] // ros ro
}
func minDistance(word1 string, word2 string) int {
	/*
		note！如果有两个string 入参，dp的状态定义一般都是二维，前ij各表示一个string
		dp[i][j] : word1前i个字符 和 word2前j个字符的最小操作次数

	*/
	//dp := [][]int{}
	dp := Mem{}

	for i := 0; i <= len(word1); i++ {
		dpi := make([]int, len(word2)+1)
		dp = append(dp, dpi)
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	//
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min3(dp.Delete(i, j), dp.Insert(i, j), dp.Repalce(i, j))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
func minimumTotal(triangle [][]int) int {
	cur := []int{}
	pre := []int{}
	for i, e := range triangle {
		pi := i - 1
		if pi < 0 {
			cur = append(cur, e[0])
		} else {
			for ei, ee := range e {
				if ei-1 < 0 {
					cur = append(cur, pre[ei]+ee)
				} else if ei == len(e)-1 {
					cur = append(cur, pre[ei-1]+ee)
				} else {
					cur = append(cur, min2(pre[ei], pre[ei-1])+ee)
				}
			}
		}
		pre = cur
		cur = make([]int, 0)
	}
	min := 9999999999
	for _, ei := range pre {
		min = min2(min, ei)
	}
	return min
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3(a, b, c int) int {
	t := min2(a, b)
	return min2(t, c)
}
