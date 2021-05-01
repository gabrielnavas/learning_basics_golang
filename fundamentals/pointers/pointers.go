package main

import "fmt"

func main() {
	var age = 22

	var ageP *int = nil // criei variavel
	ageP = &age         //atribui para o endereco de age

	fmt.Println(ageP)  // 0xc000018110
	fmt.Println(*ageP) // 22
	age++
	*ageP++
	fmt.Println(*ageP) // 24
	fmt.Println(age)   // 24

	// ageP++ isso não da certo.
}
