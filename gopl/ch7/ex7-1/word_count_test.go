package ch7

import (
	"fmt"
	"testing"
)

func TestWordCount(t *testing.T) {
	var wc WordCounter
	wc.write([]byte("hello world! \n123"))
	fmt.Println(wc.Word)
	fmt.Println(wc.Line)
}
