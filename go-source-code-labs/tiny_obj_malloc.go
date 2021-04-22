package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"

	//"runtime"
	//"unsafe"
)

func main()  {
	fnc()
}
type myStruct struct {
	P [1<<22]int64
}
//go:noinline
func fnc()  {
	mx := &myStruct{} // mx -> 0xc00001c080
	my := &myStruct{}  // my -> 0xc00001c090
	mx.P[0]= 1000
	my.P[0]= 2
	fmt.Println(unsafe.Sizeof(*mx))
	oldmx := uintptr(unsafe.Pointer(mx))
	mx = my
	for i:=0;i<100;i++ {
		runtime.GC()
		debug.FreeOSMemory()
	}

	oldmxStc := (*myStruct)(unsafe.Pointer(oldmx))
	fmt.Println(oldmxStc.P[0])
}

