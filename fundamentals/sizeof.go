package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = 1
	var b byte = 1
	var c int16 = 1
	var d int32 = 1
	var e int64 = 1
	var f int = 1

	fmt.Println(unsafe.Sizeof(a)) // 1
	fmt.Println(unsafe.Sizeof(b)) // 1
	fmt.Println(unsafe.Sizeof(c)) // 2
	fmt.Println(unsafe.Sizeof(d)) // 4
	fmt.Println(unsafe.Sizeof(e)) // 8
	fmt.Println(unsafe.Sizeof(f)) // 8
}
