package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// unsigned byte(uint8), uint8, uint16, uint32, uint64
	var b byte = 255
	var uinteger_8b uint8 = math.MaxUint8
	var uinteger_16b uint16 = math.MaxUint16
	var uinteger_32b uint32 = math.MaxUint32
	var uinteger_64b uint64 = math.MaxUint64

	// type prints
	fmt.Println(reflect.TypeOf(b))            // uint8
	fmt.Println(reflect.TypeOf(uinteger_8b))  // uint8
	fmt.Println(reflect.TypeOf(uinteger_16b)) // uint16
	fmt.Println(reflect.TypeOf(uinteger_32b)) // uint32
	fmt.Println(reflect.TypeOf(uinteger_64b)) // uint64

	// map of the Unicode table  (int32)
	var int32Alias rune = 'a'

	// int8, int16, int32, int64
	var integer_8 int8 = math.MaxInt8
	var integer_16 int8 = math.MaxInt8
	var integer_32 int8 = math.MaxInt8
	var integer_64 int8 = math.MaxInt8

	fmt.Println(int32Alias)                 // code 97 letter 'a' from unicode table
	fmt.Println(reflect.TypeOf(int32Alias)) // int32
	fmt.Println(reflect.TypeOf(integer_8))  // int8
	fmt.Println(reflect.TypeOf(integer_16)) // int16
	fmt.Println(reflect.TypeOf(integer_32)) // int32
	fmt.Println(reflect.TypeOf(integer_64)) // int64
	fmt.Println(reflect.TypeOf(123))        // int (default)

	// float32 and float64
	var float_32 float32 = math.MaxFloat32
	var float_64 float64 = math.MaxFloat64

	fmt.Println(reflect.TypeOf(float_32)) // float32
	fmt.Println(reflect.TypeOf(float_64)) // float64
	fmt.Println(reflect.TypeOf(48.99))    // float64 (default)

	// bool
	var boo_true bool = true
	var boo_false bool = false
	fmt.Println(reflect.TypeOf(boo_true))  // bool
	fmt.Println(reflect.TypeOf(boo_false)) // bool

	// string
	var name string = "Gabriel"
	fmt.Println(reflect.TypeOf(name)) // string
	fmt.Println(len(name))            // 7

	// multiple string lines
	var description string = `
		My 
		name
		is
		Gabriel!! 
	`
	fmt.Println(len(description))

	// char??? dont have, char is int32
	letter := 'g'
	var otherChar int32 = 'a'
	letterIsEqualsOtherChar := reflect.TypeOf(letter) == reflect.TypeOf(otherChar)
	fmt.Println(letterIsEqualsOtherChar) // true
}
