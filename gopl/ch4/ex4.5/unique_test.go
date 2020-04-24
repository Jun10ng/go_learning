package ex4_5

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T)  {
	s := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e","e"}
	ans := unique(s)
	fmt.Println(ans)
}
