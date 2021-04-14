package main

import (
	"fmt"
	"math"
	"math/rand"
)

// math.Round(x*100)不会丢失精度
func main() {
	for i := 0; i < 10000; i++ {
		x := float64(rand.Intn(10000)) / 100
		if (int64)(x*100) != (int64)(math.Round(x*100)) {
			fmt.Printf("x:%v, x*100:%v, round(x*100):%v\n", x, (int64)(x*100), (int64)(math.Round(x*100)))
		}
	}

}
