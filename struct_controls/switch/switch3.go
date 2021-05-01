package main

import (
	"fmt"
	"time"
)

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "unknow"
	}
}

func main() {
	fmt.Println(getType(2))          //int
	fmt.Println(getType(5.22))       //float
	fmt.Println(getType("Gabriel"))  //string
	fmt.Println(getType(func() {}))  //function
	fmt.Println(getType(time.Now())) //unknow
}
