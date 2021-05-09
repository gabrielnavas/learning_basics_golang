package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.\n")
	fmt.Println("Nova Linha")

	height := 1.87

	// fmt.Sprint is to transform number or float to string
	heightS := fmt.Sprint(height)
	fmt.Println(heightS)
	fmt.Printf("My height is %.2f\n", height)

	// types int, float, bool, string
	name := "Gabriel"
	age := 26
	weight := 80.00
	isMarried := false
	fmt.Printf("\n\n")
	fmt.Printf("%s %d %.2f is married %t\n", name, age, weight, isMarried)

	// v is to all types
	fmt.Printf("%v %v %v2f is married %v\n", name, age, weight, isMarried)

	var a int8 = 21
	var b byte = 21
	var c int16 = 21
	var d int32 = 21
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
}
