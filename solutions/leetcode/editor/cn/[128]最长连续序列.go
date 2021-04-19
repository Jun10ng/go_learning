package main
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。 
//
// 
//
// 进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？ 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。 
//
// 示例 2： 
//
// 
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 104 
// -109 <= nums[i] <= 109 
// 
// Related Topics 并查集 数组 
// 👍 726 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func longestConsecutive(nums []int) int {
	u := NewUnionFindSet(nums)
	pre := make(map[int]int,len(nums))
	for i,e := range nums {
		if _,exist := pre[e];exist {
			continue
		}
		// 连续，找e+1和e-1
		if prei,ok := pre[e+1];ok{
			u.Union(i,prei)
		}
		if prei,ok := pre[e-1];ok{
			u.Union(i,prei)
		}
		pre[e] = i
	}
	return u.LC
}

type UnionFindSet struct {
	M int
	Nums []int
	Parent []int
	Rank   []int
	LC int  // 最大长度
}
func NewUnionFindSet(nums []int) *UnionFindSet {
	m := len(nums)
	if m==0 {
		return &UnionFindSet{}
	}

	p := make([]int,m)
	r := make([]int,m)
	for i,_:=range p{
		p[i] = i //每个元素的父亲都是自己
		r[i] = 1
	}
	return &UnionFindSet{
		M: m,
		Nums: nums,
		Parent: p,
		Rank: r,
		LC: 1,
	}
}
func (u *UnionFindSet)find(x int) int {
	px := u.Parent[x]
	if u.Parent[x] == x{ //找到父节点
		return x
	}
	return u.find(px)
}
func (u *UnionFindSet)Union(x,y int)  {
	px := u.find(x)
	py := u.find(y)
	if px != py{
		if u.Nums[x] > u.Nums[y]{
			u.Parent[x] = py
			u.Rank[py]+=u.Rank[x]
			u.LC = max2(u.Rank[py],u.LC)
		}else {
			u.Parent[y] = px
			u.Rank[px]+=u.Rank[y]
			u.LC = max2(u.Rank[px],u.LC)
		}
	}
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
