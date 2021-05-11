package main

import "fmt"

// executing using: go run *.go
func main() {
	p1 := Pointer{2.0, 2.0}
	p2 := Pointer{2.0, 4.0}

	cx, cy := catetos(p1, p2)
	fmt.Println("cx, cy: ", cx, cy)

	distance := Distance(p1, p2)
	fmt.Println("distance: ", distance)
}
