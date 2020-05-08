package ch7

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}
	c, n := CountingWriter(b)
	data := []byte("hi there")
	c.Write(data)
	fmt.Println(*n)
}
