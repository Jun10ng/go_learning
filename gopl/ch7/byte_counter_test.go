package ch7

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello,%s", name)
	fmt.Println(c)
}


