package main

import (
	"fmt"
	"time"
)

/*
switch true {
	case hour < 12:
		return "good morning"
	case hour < 18:
		return "good afternoon"
	default:
		return "good night"
	}
*/
func getPeriod(hour int) string {
	switch {
	case hour < 12:
		return "good morning"
	case hour < 18:
		return "good afternoon"
	default:
		return "good night"
	}
}

func main() {
	t := time.Now()
	period := getPeriod(t.Hour())
	fmt.Println(period)
}
