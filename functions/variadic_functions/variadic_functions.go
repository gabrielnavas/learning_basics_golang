package main

import "fmt"

func avarage(values ...float64) float64 {
	sum := 0.0
	for _, n := range values {
		sum += n
	}
	return sum / float64(len(values))
}

func main() {
	n := avarage(2, 4, 6, 8)
	fmt.Printf("avarage is: %v\n", n)
}
