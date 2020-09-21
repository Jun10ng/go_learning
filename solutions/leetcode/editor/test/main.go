package main

import "fmt"

func main()  {
	fmt.Print(partition("aba"))
	//fmt.Print(ok("aabaa"))
}
func partition(s string) [][]string {
	ans := [][]string{}
	path := []string{}
	trace := func(path []string, idx int) {}

	trace = func(path []string,idx int) {
		if idx == len(s){
			cp := make([]string,len(path))
			copy(cp,path)
			ans = append(ans,cp)
			return
		}
		for i:= idx;i<len() {
			
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