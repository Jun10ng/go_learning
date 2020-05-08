package ch7

import (
	"fmt"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "123456"
	input := strings.NewReader(s)
	lr := LimitReader(input, 5)
	n, err := lr.Read([]byte(s))
	fmt.Println(n, err)
}
