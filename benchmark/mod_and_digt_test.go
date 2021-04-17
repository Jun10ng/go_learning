package benchmark

import (
	"testing"
)

func BenchmarkDigts(b *testing.B) {
	for i := 0; i <= 9000000000; i++ {
		bit(i)
	}
}
func bit(i int) bool {
	return i&1 == 0
}
func BenchmarkMod(b *testing.B) {
	for i := 0; i <= 9000000000; i++ {
		mod(i)
	}
}
func mod(i int) bool {
	return i%2 == 0
}
