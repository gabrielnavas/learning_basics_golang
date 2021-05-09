package main

import "fmt"

func main() {
	// go vet fundamentals/commands.go
	// fundamentals/commands.go:6:2: Printf format %s reads arg #1, but call has 0 args

	fmt.Printf("hello %s")
}
