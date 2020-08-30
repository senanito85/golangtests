package main

import (
	"fmt"
)

var min int = 1
var max int = 9

func main() {
	fmt.Println("")
	secondary(min)
	tertiery(max)
}

func secondary(y int) {
	if y <= max {
		fmt.Printf("%d  - ", y)

		for i := 1; i <= y; i++ {
			fmt.Printf("*")
		}

		fmt.Println("")
		y++
		secondary(y)

	} else {
		fmt.Println("")
	}
}

func tertiery(x int) {
	if x > 0 {
		fmt.Printf("%d  - ", x)

		for i := x; i > 0; i-- {
			fmt.Printf("*")
		}

		fmt.Println("")
		x--
		tertiery(x)

	} else {
		fmt.Println("")
	}
}
