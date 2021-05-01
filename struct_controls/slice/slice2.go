package main

import "fmt"

func main() {
	houses := []int{55, 77, 100, 22, 88}
	fmt.Println(houses[:2]) // [55, 77]
	fmt.Println(houses[1:]) // [77, 100, 22, 88]
}
