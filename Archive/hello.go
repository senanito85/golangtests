package main

import (
	//  "os"
	"fmt"
	"strconv"
)

type point struct {
	x, y int
}

func main() {
	name := "Sanan"
	age := 34
	fmt.Printf("Hello %s you are %d\n", name, age)

	var mystring = "1"
	x := 10
	number, _ := strconv.Atoi(mystring)
	fmt.Println(number)
	fmt.Println(x)
	x = number
	fmt.Println(x)
}
