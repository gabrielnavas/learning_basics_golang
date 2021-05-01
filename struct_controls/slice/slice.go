package main

import (
	"fmt"
	"reflect"
)

func main() {
	scores := [3]int{2, 4, 6}           //array
	ages := []int{1, 2, 3}              //slice
	fmt.Println(reflect.TypeOf(scores)) // [3]int
	fmt.Println(reflect.TypeOf(ages))   // []int
}
