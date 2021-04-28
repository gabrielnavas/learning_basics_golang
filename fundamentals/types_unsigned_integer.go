package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// unsigned byte == uint8, uint16, uint32, uint64
	// unsigned byte
	var b byte = 255

	// unsigned uint8
	var integer_8b uint8 = 255

	// unsigned uint16
	var integer_16b uint16 = uint16(math.Pow(float64(integer_8b), 2))

	// unsigned uint32
	var integer_32b uint32 = uint32(math.Pow(float64(integer_16b), 2))

	// unsigned uint64
	var integer_64b uint64 = uint64(math.Pow(float64(integer_32b), 2))

	// type prints
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(integer_8b))
	fmt.Println(reflect.TypeOf(integer_16b))
	fmt.Println(reflect.TypeOf(integer_32b))
	fmt.Println(reflect.TypeOf(integer_64b))
}
