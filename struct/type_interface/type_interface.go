package main

import "fmt"

type Car struct {
	Name string
}

func main() {
	var anythink interface{}

	anythink = 2
	fmt.Println(anythink) // 2

	anythink = "Gabriel"
	fmt.Println(anythink) // Gabriel

	anythink = false
	fmt.Println(anythink) // false

	anythink = Car{
		Name: "Gabriel",
	}
	fmt.Println(anythink) // {Gabriel}
}
