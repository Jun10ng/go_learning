package ex4_7

import (
	"fmt"
	"testing"
)

func TestRevInPlace(t *testing.T)  {
	b := []byte("abcde")
	ans:=reverseInPlace(b)
	fmt.Println(string(ans))
}