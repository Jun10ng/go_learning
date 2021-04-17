package main

func main() {
	c := createChan()
	xx := "123123"
	fnca(c, &xx)
	x := fncb(c)
	println(x)
}

//go:noinline
func createChan() chan string {
	cc := make(chan string, 1)
	return cc
}

//go:noinline
func fnca(c chan string, xptr *string) {
	cx := *xptr
	c <- cx
}

//go:noinline
func fncb(c chan string) string {
	return <-c
}
