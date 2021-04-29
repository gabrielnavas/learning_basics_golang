package main

import "fmt"

/*
	Fundamentos-de-Matematica-Elementar-Volume-1-Conjuntos-e-Funcoes
	A.12 - a
	conjunto dos multiplos inteiros de 3,
	entre -10 e +10
*/
func A12a() []int {
	const lessTen = -9
	const max = 10
	set := make([]int, 0)
	for i := lessTen; i < max; i += 3 {
		set = append(set, i)
	}
	return set
}

/*
	Fundamentos-de-Matematica-Elementar-Volume-1-Conjuntos-e-Funcoes
	A12 - b
*/
func A12b() []int {
	const init_one = 1
	const fourty_two = 42
	set := make([]int, 0)
	for i := init_one; i <= fourty_two; i++ {
		if fourty_two%i == 0 {
			set = append(set, i)
		}
	}
	return set
}

func main() {
	fmt.Printf("A12a: %v\n", A12a())
	fmt.Printf("A12b: %v\n", A12b())
}
