package main
//
// 
// 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连
//。 
//
// 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。 
//
// 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 
//isConnected[i][j] = 0 表示二者不直接相连。 
//
// 返回矩阵中 省份 的数量。 
//
// 
//
// 示例 1： 
//
// 
//输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
//输出：3
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 200 
// n == isConnected.length 
// n == isConnected[i].length 
// isConnected[i][j] 为 1 或 0 
// isConnected[i][i] == 1 
// isConnected[i][j] == isConnected[j][i] 
// 
// 
// 
// Related Topics 深度优先搜索 并查集 
// 👍 521 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func findCircleNum(isConnected [][]int) int {
	u := NewUnionFindSet(isConnected)
	for i:=0;i<u.M;i++{
		for j:=i;j<u.M;j++{
			if isConnected[i][j]==1{
					u.Union(i,j)
			}
		}
	}
	return u.Cnt
}

type UnionFindSet struct {
	M, N, Cnt int
	Grid      [][]int
	Parent    []int
	Rank      []int //省份内城市大小个数
}

func NewUnionFindSet(g [][]int) *UnionFindSet {
	m := len(g)
	n := len(g[0])
	cnt := 0
	if n == 0 {
		return &UnionFindSet{}
	}
	p := make([]int, m)
	r := make([]int, m)
	for i, _ := range p {
		p[i] = i
		cnt++
	}

	return &UnionFindSet{
		Grid:   g,
		Parent: p,
		Rank:   r,
		M:      m,
		N:      n,
		Cnt:    cnt,
	}
}
func (u *UnionFindSet) Find(index int) int {
	if u.Parent[index] == index { //找到根节点
		return index
	}
	return u.Find(u.Parent[index]) // 递归子节点
}

func (u *UnionFindSet) Union(x, y int) {
	rootx := u.Find(x)
	rooty := u.Find(y)
	if rootx != rooty {
		//两个节点 根节点不一样，
		//但是两个节点相邻，所以进行union,cnt--
		//把深度(rank)小的一方指向深度大的
		if u.Rank[rootx] > u.Rank[rooty] {
			u.Parent[rooty] = rootx
		} else if u.Rank[rooty] > u.Rank[rootx] {
			u.Parent[rootx] = rooty
		} else {
			// 两个节点深度相同，选择让rooty指向rootx，rootx深度+1
			u.Parent[rooty] = rootx
			u.Rank[rootx]+= u.Rank[rooty]
		}
		u.Cnt--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
