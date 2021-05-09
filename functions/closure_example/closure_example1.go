package main

import (
	"fmt"
	"strings"
)

func example(name string) func() {
	return func() {
		fmt.Println(name)
	}
}

func example2() func(fullName string) {
	return func(fullName string) {
		splited := strings.Split(fullName, " ")
		first, last := splited[0], splited[1]
		fmt.Println(first, last)
	}
}

func main() {
	name := "Gabriel"
	printName := example(name)
	printName()
}
