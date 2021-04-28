package main

import (
	"fmt"
	"strconv"
)

func main() {
	const price = 1000.00
	const parts = 12
	const bitPrice = price / parts
	fmt.Println(bitPrice) // 83.33333333333333

	/*
		warning, string(number) == ascii table
		97 == a
	*/
	myAge := "My age is " + string(97)
	fmt.Println(myAge)

	// int to string
	myNumber := 10
	fmt.Println("my number is " + strconv.Itoa(myNumber)) // my number is 10

	binary := 2
	octal := 8
	hex := 16
	fmt.Println("10 in binary is: " + strconv.FormatInt(int64(myNumber), binary))
	fmt.Println("10 in octal is: " + strconv.FormatInt(int64(myNumber), octal))
	fmt.Println("10 in hex is: " + strconv.FormatInt(int64(myNumber), hex))

	// string to int
	num, _ := strconv.Atoi("10")

	// num to string again (haha)
	fmt.Println("convert 10 to: " + strconv.Itoa(num))

	// string boolean to boolean
	go_lang_is_nice, _ := strconv.ParseBool("true")
	fmt.Println(go_lang_is_nice) // true
	if go_lang_is_nice {
		fmt.Println("Yeah, it is nice!")
	}

	// other value to boolean
	any_value_alfa, _ := strconv.ParseBool("any_value_alfa")
	less123, _ := strconv.ParseBool("123")
	more123, _ := strconv.ParseBool("-123")
	zero, _ := strconv.ParseBool("0")
	one, _ := strconv.ParseBool("1")
	fmt.Println(any_value_alfa) // false
	fmt.Println(less123)        // false
	fmt.Println(more123)        // false
	fmt.Println(zero)           // false
	fmt.Println(one)            // true

}
