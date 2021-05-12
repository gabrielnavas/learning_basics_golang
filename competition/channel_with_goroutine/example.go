package main

import (
	"fmt"
	"time"
)

// Channel is the form of communication between goroutines
// Channle is a type

func twoThreeFourX(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //send data sync

	time.Sleep(time.Second)
	c <- 3 * base //send data sync

	time.Sleep(3 * time.Second)
	c <- 4 * base //send data sync
}

func main() {
	c := make(chan int)
	go twoThreeFourX(2, c)

	firstC := <-c
	secondC := <-c
	fmt.Println(firstC, secondC)
	fmt.Println(<-c)
}
