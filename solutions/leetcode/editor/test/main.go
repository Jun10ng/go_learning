package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(helper("1232"))
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
