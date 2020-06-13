package ex4_3

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	origin := [...]int{0, 1, 2, 3, 4}
	reverse(&origin)
	fmt.Println(origin)
}
