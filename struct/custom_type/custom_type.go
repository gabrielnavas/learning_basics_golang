package main

import "fmt"

type Score float64

func (n Score) Between(init, end float64) bool {
	if n >= Score(init) && n <= Score(end) {
		return true
	}
	return false
}

func generateScore(score Score) string {
	if score.Between(9.0, 10.0) {
		return "A"
	}
	if score.Between(6.0, 8.9) {
		return "B"
	}
	if score.Between(3.0, 5.9) {
		return "C"
	}
	return "D"
}

func main() {
	var score1 Score = 9.8
	var score2 Score = 6.8
	var score3 Score = 3.8
	var score4 Score = 2.8
	fmt.Println(generateScore(score1)) // A
	fmt.Println(generateScore(score2)) // B
	fmt.Println(generateScore(score3)) // C
	fmt.Println(generateScore(score4)) // D
}
