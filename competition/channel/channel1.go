package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // send data to channel
	<-ch    // read data from channel

	ch <- 2           // send other data to channel
	fmt.Println(<-ch) //read data from channel and print
}
