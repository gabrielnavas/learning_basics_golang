package main

import (
	"fmt"
	"strings"
)

func example() func(fullName string) {
	return func(fullName string) {
		splited := strings.Split(fullName, " ")
		first, last := splited[0], splited[1]
		fmt.Println(first, last)
	}
}

func main() {
	printNames := example()
	printNames("Gabriel Silva")
}
