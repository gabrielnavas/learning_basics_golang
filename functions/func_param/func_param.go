package main

import "fmt"

func sum(tot, value int) (totMoreValue int) {
	totMoreValue = tot + value
	return
}

func calc(sumCalc func(tot, value int) (totMoreValue int), values []int) (sum int) {
	for _, n := range values {
		sum += sumCalc(sum, n)
	}
	return
}

func main() {
	values := []int{2, 4, 6, 8, 10}
	sumValues := calc(sum, values)
	fmt.Printf("sum of the: %v is %v\n", values, sumValues)
}
