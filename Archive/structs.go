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

func main() {
	fmt.Println("Helloe \n")

	//method one
	d := dealerShip{name: "A1 Auto", city: "Melbourne"}
	fmt.Println("City = " + d.city + "\nDealership = " + d.name + "\n")

	//method two
	var d2 dealerShip
	d2 = dealerShip{name: "Discount Auto", city: "Perth"}
	fmt.Println("City = " + d2.city + "\nDealership = " + d2.name + "\n")

	//methid 3
	d3 := dealerShip{}
	d3.city = "Sydney"
	d3.name = "Luxary Auto"
	fmt.Println("City = " + d3.city + "\nDealership = " + d3.name + "\n")
}
