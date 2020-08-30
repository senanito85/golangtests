package main

import (
	//  "os"
	"fmt"
)

//creating a struct
type dealerShip struct {
	name string
	city string
}

func calculateenvUtil(remaining float32, start float32, dealer dealerShip) (percentSold float32, dealerName string) {
	dealerName = dealer.name
	percentSold = 1 - remaining/start
	return

}

func main() {

	fmt.Println("\n------------------------------------------------------------")
	fmt.Println("|    City     |     Dealership     |       Sold Cars       |")
	fmt.Println("------------------------------------------------------------")

	//method one
	d := dealerShip{name: "A1 Auto", city: "Melbourne"}
	fmt.Print(d.city)
	sold, name := calculateenvUtil(149, 200, d)
	fmt.Printf("          "+name+"              "+"%-1.2f\n", sold)

	//method two
	var d2 dealerShip
	d2 = dealerShip{name: "Discount Auto", city: "Perth"}
	fmt.Print(d2.city)
	sold2, name2 := calculateenvUtil(99, 200, d2)
	fmt.Printf("              "+name2+"        "+"%-1.2f\n", sold2)

	//methid 3
	d3 := dealerShip{}
	d3.city = "Sydney"
	d3.name = "Luxary Auto"
	fmt.Print(d3.city)
	sold3, name3 := calculateenvUtil(175, 200, d3)
	fmt.Printf("             "+name3+"          "+"%-1.2f\n", sold3)
	fmt.Println("------------------------------------------------------------")

}
