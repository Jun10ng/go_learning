package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// Defining a struct type L
type L struct {
	x, y int
	s    string
}
type L2 struct {
	x    int
	y    int
	s, b string
}

func main() {
	var PL *L
	// Defining *addr unsafe.Pointer
	var unsafepL = (*unsafe.Pointer)(unsafe.Pointer(&PL))

	// Defining value
	// of unsafe.Pointer
	var px = L{
		1, 2, "s",
	}

	atomic.StorePointer(unsafepL, unsafe.Pointer(&px))

	// Printed if value is stored
	fmt.Println(PL)
	fmt.Println("Val Stored!")

	// 两个结构一样的结构体
	var PL2 *L2
	pl2 := (*unsafe.Pointer)(unsafe.Pointer(&PL2))
	atomic.StorePointer(pl2, unsafe.Pointer(&px))
	fmt.Println(PL2)
}
