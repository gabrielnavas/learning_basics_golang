package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(1 == 1)   //true
	fmt.Println(2 > 1)    // true
	fmt.Println(10 < 5)   //false
	fmt.Println(5 <= 10)  //true
	fmt.Println(20 >= 10) //true

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println(d1 == d2) // true

	type Person struct {
		name string
	}

	my := Person{"Gabriel"}
	aunt := Person{"Maria"}
	fmt.Println(my == aunt) // false
	fmt.Println(my == my)   //true
}
