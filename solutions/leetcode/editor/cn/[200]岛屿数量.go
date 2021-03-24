package main
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。 
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。 
//
// 此外，你可以假设该网格的四条边均被水包围。 
//
// 
//
// 示例 1： 
//
// 
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 300 
// grid[i][j] 的值为 '0' 或 '1' 
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集 
// 👍 1050 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func numIslands(grid [][]byte) int {
	u := NewUnionFindSet(grid)
	for i,ei:=range grid{
		for j,ej:=range ei{
			if ej=='1'{
				for _,v :=range [][]int{{-1,0},{1,0},{0,-1},{0,1}}{
					vi,vj := i+v[0],j+v[1]
					if vi >= 0 && vj >= 0 && vi < u.M && vj < u.N && grid[vi][vj]=='1' {
						u.Union(i*u.N+j,vi*u.N+vj)
					}
				}
			}
		}
	}
	return u.Cnt
}
type UnionFindSet struct {
	M,N,Cnt int
	Grid [][]byte
	Parent []int
	Rank []int //深度
}

func NewUnionFindSet(g [][]byte) *UnionFindSet {
	m:=len(g)
	n:=len(g[0])
	cnt := 0
	if n==0{
		return &UnionFindSet{}
	}
	p := make([]int,m*n)
	r := make([]int,m*n)
	for i,_:=range p{
		p[i] = -1
	}
	for i,ei:=range g{
		for j,ej:=range ei{
			if ej=='1'{
				p[i*n+j]=i*n+j
				cnt++
			}
		}
	}
	return &UnionFindSet{
		Grid: g,
		Parent: p,
		Rank: r,
		M: m,
		N: n,
		Cnt: cnt,
	}
}
func (u *UnionFindSet)Find(index int)int{
	if u.Parent[index]==index {//找到根节点
		return index
	}
	return u.Find(u.Parent[index]) // 递归子节点
}

func (u *UnionFindSet)Union(x,y int)  {
	rootx:=u.Find(x)
	rooty:=u.Find(y)
	if rootx != rooty {
		//两个节点 根节点不一样，
		//但是两个节点相邻，所以进行union,cnt--
		//把深度(rank)小的一方指向深度大的
		if u.Rank[rootx]>u.Rank[rooty] {
			u.Parent[rooty] = rootx
		}else if u.Rank[rooty]>u.Rank[rootx] {
			u.Parent[rootx] = rooty
		}else {
			// 两个节点深度相同，选择让rooty指向rootx，rootx深度+1
			u.Parent[rooty] =rootx
			u.Rank[rootx]++
		}
		u.Cnt--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
