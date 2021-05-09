package main

import "fmt"

func main() {
	printAnything()
	printName("Gabriel")
	printNameAndLastName("Gabriel", "Navas")
	name, age := getInfo()
	name2, _ := getInfo()
	fmt.Println(name, age)
	fmt.Println(name2)
}

func printAnything() {
	fmt.Println("printAnything")
}

func printName(name string) {
	fmt.Println(name)
}

func printNameAndLastName(firstName, lastName string) {
	fmt.Println("Mr." + lastName + ", " + firstName)
}

func getInfo() (string, int8) {
	return "Gabriel", 25
}
