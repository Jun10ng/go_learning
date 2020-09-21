package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(isAdditiveNumber("1235813"))
	//fmt.Print(ok("aabaa"))
}
func isAdditiveNumber(num string) bool {
	ret := false
	//l,ml := len(num),len(num)/2 //单个数字长度不大于ml
	trace := func(cnt, n1, n2 int, rest string) {}
	// cnt 是目前池内数字，当cnt > 2 才会判断
	// rest 是可选择的string

	trace = func(cnt, n1, n2 int, rest string) {
		if len(rest) == 0 {
			if cnt > 2 {
				ret = true
			}
			return
		}

		//开始从rest中取值，添加分支并剪枝
		n3 := ""
		for _, e := range rest {
			n3 = n3 + string(e)
			// 0 可以新建分支 01 02 则剪枝
			if n3[0] == '0' && len(n3) != 1 {
				return
			}
			e, _ := strconv.Atoi(n3)
			if cnt >= 2 {
				if e == n1+n2 {
					trace(cnt+1, n2, e, rest[len(n3):])
				}
			} else {
				trace(cnt+1, n2, e, rest[len(n3):])
			}
			if ret {
				return
			}
		}
	}

	//199100199
	//
	trace(0, 0, 0, num)
	return ret
}
