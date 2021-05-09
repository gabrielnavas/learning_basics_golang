package main

import "fmt"

func changeValues(x, y int) (ty, tx int) {
	tx = x
	ty = y
	return // return clear
}

func main() {
	x := 2
	y := 4

	valueX, valueY := changeValues(2, 4)
	fmt.Printf("x: %d y: %d valueX: %d valeY: %d\n", x, y, valueX, valueY)
}
