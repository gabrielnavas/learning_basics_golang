package main

import (
	"fmt"
)

func fat(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("N is %v", n)
	case n == 0 || n == 1:
		return n, nil
	default:
		nLessOne, _ := fat(n - 1)
		return n * nLessOne, nil
	}
}

func main() {
	n1, error1 := fat(-1)
	if error1 != nil {
		fmt.Println(error1)
	} else {
		fmt.Println(n1)
	}

	n2, error2 := fat(0)
	if error2 != nil {
		fmt.Println(error1)
	} else {
		fmt.Println(n2)
	}

	n3, error3 := fat(1)
	if error3 != nil {
		fmt.Println(error3)
	} else {
		fmt.Println(n3)
	}

	n4, error4 := fat(4)
	if error4 != nil {
		fmt.Println(error4)
	} else {
		fmt.Println(n4)
	}
}
