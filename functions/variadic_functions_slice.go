package main

import "fmt"

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	names := []string{"Gabriel", "Anna", "John", "Richard"}
	printNames(names...)
}
