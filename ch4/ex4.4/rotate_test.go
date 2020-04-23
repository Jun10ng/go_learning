package ex4_4

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T)  {
	s := []int{1,2,3,4,5}
	n := 2
	s=rotate(s,n)
	fmt.Println(s)
}
