package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
	//"runtime"
	//"unsafe"
)

func main() {
	fnc()
}

type myStruct struct {
	P [1 << 20]int64
}

//go:noinline
func fnc() {
	mx := &myStruct{}
	my := &myStruct{}
	mx.P[0] = 1
	my.P[0] = 2

	oldmx := uintptr(unsafe.Pointer(mx))

	// 原本的mx, tx 已经没有被引用了
	mx = my

	for i := 0; i < 100; i++ {
		runtime.GC()
		debug.FreeOSMemory()
	}

	oldmxStc := (*myStruct)(unsafe.Pointer(oldmx))
	//oldtxStc := (*myTiny)(unsafe.Pointer(oldtx))
	fmt.Println(oldmxStc.P[0])
	//fmt.Println(oldtxStc.P[0])
}
