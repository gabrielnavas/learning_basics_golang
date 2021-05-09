package main

import (
	"fmt"
)

/*
	uint accepts only values equal to or greater than 0,
	there is no need to check if it is less than 0
*/
func fat(n uint) uint {
	switch n {
	case 0, 1:
		return 1
	}
	return n * fat(n-1)
}

func main() {
	n1 := fat(0)
	n2 := fat(1)
	n3 := fat(4)
	fmt.Println(n1) // 1
	fmt.Println(n2) // 1
	fmt.Println(n3) // 24
}
