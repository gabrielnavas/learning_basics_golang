package main

import (
	"fmt"
	"math"
)

func average(scores [3]float64) float64 {
	if len(scores) == 0 {
		return math.NaN()
	}
	sum := 0.0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	average := sum / float64(len(scores))
	return average
}

func main() {
	var scores [3]float64
	scores[0], scores[1], scores[2] = 3, 1, 2
	average := average(scores)
	fmt.Println(average)
}
