package ch7

import (
	"io"
)

type byteCounter struct {
	w       io.Writer
	written int64
}

func (bc *byteCounter) Write(p []byte) (int, error) {
	n, err := bc.w.Write(p)
	bc.written += int64(n)
	return n, err
}

// CountingWriter return a pointer，after calling this func，the value of pointer referened will be change
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := &byteCounter{w, 0}
	return bc, &bc.written
}
