package ex4_6

import (
	"fmt"
	"testing"
)

func TestReplaceSpace(t *testing.T)  {
	b := []byte("a b   c     d")
	ans :=replaceSpace(b)
	fmt.Printf("[]byte is %v\n",ans)
	fmt.Printf("to []string is %v\n",string(ans))
}