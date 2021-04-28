package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// unsigned byte
	var b byte = 255

	// unsigned uint8
	// 255
	var integer_8b uint8 = math.MaxUint8

	// unsigned uint16
	// 65535
	var integer_16b uint16 = math.MaxUint16

	// unsigned uint32
	// 4294967295
	var integer_32b uint32 = math.MaxUint32

	// map of the Unicode table  (int32)
	var int32Alias rune = 'a'
	fmt.Println(int32Alias) // code 97 letter 'a' from unicode table

	// unsigned uint64
	// 18446744073709551615
	var integer_64b uint64 = math.MaxUint64

	// type prints
	fmt.Println(reflect.TypeOf(b))           // uint8
	fmt.Println(reflect.TypeOf(integer_8b))  // uint8
	fmt.Println(reflect.TypeOf(integer_16b)) // uint16
	fmt.Println(reflect.TypeOf(integer_32b)) // uint32
	fmt.Println(reflect.TypeOf(int32Alias))  // uint32
	fmt.Println(reflect.TypeOf(integer_64b)) // uint64
}
