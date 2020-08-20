package main

import (
	"container/list"
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)

func main() {
	stack := list.New()
	stack.PushBack(1)
	stack.PushBack(2)

	stack.Remove(stack.Back())
	fmt.Print(stack.Back())
}
