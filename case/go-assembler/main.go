package main

import "fmt"

func main() {

	a := 1

	//fna(a, nil)
	b := &fn{
		D: 2,
		C: 3,
		E: nil,
	}
	fmt.Println(&b)
	fmt.Println(&b.C)
	fmt.Println(&b.D)
	f := &arg{
		13,
	}
	fmt.Println(&f.F)
	b.fnb(a, 10, f)
}

type fn struct {
	C int
	D int
	E *arg
}
type arg struct {
	F int
}

func fna(a int, b *fn) {
	t := a + b.C
	fmt.Println(t)
	fmt.Println(b.D)
}

func (b *fn) fnb(a int, aa int, cc *arg) {
	t := a + b.C
	fmt.Println(t)
	fmt.Println(b.D)
	fmt.Println(b.E.F)
}
