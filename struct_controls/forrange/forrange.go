package main

import "fmt"

func main() {
	scores := [...]int{2, 4, 6, 8, 10} // range fixed in 5 for example

	for index, num := range scores {
		fmt.Printf("index: %v, num: %v\n", index, num)
	}

	// less index
	for _, num := range scores {
		fmt.Printf("num: %v\n", num)
	}

	// just index
	for index := range scores {
		fmt.Printf("index: %v\n", index)
	}
}
