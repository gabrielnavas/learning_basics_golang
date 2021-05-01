package main

import "fmt"

func fallthroughghExample(score int) {
	//example 01
	switch score {
	case 10:
		fallthrough
	case 9:
		fmt.Println("i am executing 9")
	default:
		fmt.Println("invalid")
	}
}

func normal(score int) {
	//example 01
	switch score {
	case 1:
		fmt.Println("i am executing 1")
	case 2:
		fmt.Println("i am executing 2")
	case 3:
		fmt.Println("i am executing 2")
	default:
		fmt.Println("invalid")
	}
}

func multiple(score int) {
	//example 01
	switch score {
	case 1, 2, 3:
		fmt.Println("i am executing 1 or 2 or 3")
	case 4:
		fmt.Println("i am executing 4")
	case 5:
		fmt.Println("i am executing 5")
	default:
		fmt.Println("invalid")
	}

}

func main() {

	fallthroughghExample(10)
	normal(2)
	multiple(3)
}
