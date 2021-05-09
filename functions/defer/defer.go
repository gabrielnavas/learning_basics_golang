package main

import "fmt"

func getAcceptedValue(n int) int {
	defer fmt.Println("i am late")
	if n > 1000 {
		fmt.Println("its great that 1000")
		return n
	}
	fmt.Println("its less or equals that 1000")
	return n
}

func main() {
	n := getAcceptedValue(1000)
	fmt.Println(n)
}
