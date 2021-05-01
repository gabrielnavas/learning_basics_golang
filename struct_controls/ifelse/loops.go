package main

import (
	"fmt"
	"time"
)

func whileStyle() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func forClassic() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func whileTrue() {
	i := 0
	for {
		fmt.Println("hello")
		fmt.Println("how are you??")
		fmt.Println("bye")
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
		i++
	}
}

func main() {
	whileStyle()
	forClassic()
	whileTrue()
}
