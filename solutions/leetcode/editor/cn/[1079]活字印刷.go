package main

import (
	"sort"
	"strings"
)

//你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
//
// 注意：本题中，每个活字字模只能使用一次。 
//
// 
//
// 示例 1： 
//
// 输入："AAB"
//输出：8
//解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
// 
//
// 示例 2： 
//
// 输入："AAABBC"
//输出：188
// 
//
// 
//
// 提示： 
//
// 
// 1 <= tiles.length <= 7 
// tiles 由大写英文字母组成 
// 
// Related Topics 回溯算法 
// 👍 84 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
//这题 类似90题子集，但是有排列需求，所以每层处理字符串是不同的，需要好好琢磨
func numTilePossibilities(tiles string) int {
	if tiles == "" {
		return 0
	}
	if  len(tiles) == 1 {
		return len(tiles)
	}

	ans := 0
	trace := func(rest string) {}
	trace = func(rest string) {
		// AAB
		for i:=0;i<len(rest);i++ {
			if i>0&&rest[i-1]==rest[i] { // 同层去重
				continue
			}
			ans++
			tmp := rest[:i]+rest[i+1:]//每层回溯的字符串都不一样
			trace(tmp)
		}
	}

	tmp := strings.Split(tiles,"")
	sort.Strings(tmp)
	tiles = strings.Join(tmp,"")
	trace(tiles)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
