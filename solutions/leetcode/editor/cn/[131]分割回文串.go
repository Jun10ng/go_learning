package main
//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。 
//
// 返回 s 所有可能的分割方案。 
//
// 示例: 
//
// 输入: "aab"
//输出:
//[
//  ["aa","b"],
//  ["a","a","b"]
//] 
// Related Topics 回溯算法 
// 👍 362 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	ans := [][]string{}
	path := []string{}
	trace := func(path []string, idx int) {}

	trace = func(path []string,idx int) {
		if idx == len(s)&&listOk(path){
			cp := make([]string,len(path))
			copy(cp,path)
			ans = append(ans,cp)
			return
		}else if idx == len(s)&&!listOk(path) {
			return
		}

		if idx == 0 {
			path = append(path,string(s[idx]))
			trace(path,1)
		}else {
			//加入
			path[len(path)-1] = path[len(path)-1] + string(s[idx])
			trace(path,idx+1)
			path[len(path)-1] = path[len(path)-1][:len(path[len(path)-1])-1]

			//重开一个
			path = append(path,string(s[idx]))
			trace(path,idx+1)
		}

	}
	trace(path,0)
	return ans
}

// 判断是否为回文串
func listOk(path []string) bool  {
	for _,e :=range path{
		if !ok(e) {
			return false
		}
	}
	return true
}

func ok(p string) bool {
	l := len(p)-1
	for i:=0;i<=l/2;i++ {
		if p[i]!=p[l-i] {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
