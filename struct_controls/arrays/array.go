package main

import "fmt"

func main() {
	var scores [3]float64
	fmt.Println(scores) // [0 0 0]

	scores[0], scores[1], scores[2] = 3, 1, 2
	fmt.Println(scores)    // [3 1 2]
	fmt.Println(scores[1]) // 1
}
