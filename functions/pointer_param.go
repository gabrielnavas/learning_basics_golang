package main

import "fmt"

func increment(n int) {
	n++
}

func incrementWithPointer(n *int) {
	// Go supports the use of pointers, but does not support arithmetic between them.
	*n++
}

func main() {
	n := 2
	increment(n)
	fmt.Println(n) // 2
	incrementWithPointer(&n)
	fmt.Println(n) // 3
}
