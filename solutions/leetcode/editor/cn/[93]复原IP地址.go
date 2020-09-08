package main

import (
	"strconv"
	"strings"
)

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
// 有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。 
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
// 和 "192.168@1.1" 是 无效的 IP 地址。 
//
// 
//
// 示例 1： 
//
// 输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
// 
//
// 示例 2： 
//
// 输入：s = "0000"
//输出：["0.0.0.0"]
// 
//
// 示例 3： 
//
// 输入：s = "1111"
//输出：["1.1.1.1"]
// 
//
// 示例 4： 
//
// 输入：s = "010010"
//输出：["0.10.0.10","0.100.1.0"]
// 
//
// 示例 5： 
//
// 输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 3000 
// s 仅由数字组成 
// 
// Related Topics 字符串 回溯算法 
// 👍 407 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	if s ==""{
		return []string{}
	}
	ans := []string{}
	trace := func(s string, cur []string,idx,curi int){}

	trace = func(s string,cur []string,idx,curi int) {
		icuri ,_ := strconv.Atoi(cur[curi])

		if idx >= len(s) && curi == 3 {
			cp := strings.Join(cur,".")
			ans = append(ans,cp)
			return
		}else if idx>=len(s) || len(cur[curi])>3 || icuri >255 {
			return
		}


		// 剪枝
		// 往当前curi添加
		if (curi <= 3) && !(len(cur[curi])==1 && cur[curi][0]=='0') {
			cur[curi] = cur[curi] + string(s[idx])
			trace(s,cur,idx+1,curi)
			cur[curi] = cur[curi][:len(cur[curi])-1]
		}

		//往下一个curi添加
		if curi == 3 || idx == 0 {
			return
		}
		cur[curi+1] = cur[curi+1] + string(s[idx])
		trace(s,cur,idx+1,curi+1)
		cur[curi+1] = cur[curi+1][:len(cur[curi+1])-1]
	}
	cur := []string{"","","",""}
	trace(s,cur,0,0)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
