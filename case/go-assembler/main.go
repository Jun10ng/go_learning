package main

import "fmt"

func main() {
	a := 1
	b := &arg{
		C: 2,
		D: "text",
	}
	fmt.Println(&b)
	fna(a, b)
}

type arg struct {
	C int
	D string
}

func fna(a int, b *arg) {
	t := a + b.C
	fmt.Println(t)
	fmt.Println(b.D)
}
