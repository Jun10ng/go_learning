package main

import (
	"fmt"
	"math"
)

func CalcMarkup(percent float64, price int64) int64 {
	factor := int64(10000) / 100
	percentFloat := percent / 100
	return int64(math.Ceil(float64(price/factor)*(1+percentFloat))) * factor
}

func main() {
	fmt.Println(CalcMarkup(3.5, 1230000))
}
