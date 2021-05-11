package main

import (
	"fmt"
	"time"
)

func talk(name, text string, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (count -> %d)\n", name, text, i)
	}
}

func main() {
	// talk("Gabriel", "hey you!!", 3)
	// talk("Anna", "Hello!!", 3)
	go talk("Mary", "i am Mary!!", 3)
	go talk("Jonh", "i am Mary!!", 3)
	time.Sleep(time.Second * 4)
}
