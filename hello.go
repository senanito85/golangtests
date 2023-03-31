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

	//method one.......#
	d := dealerShip{name: "A1 Auto", city: "Melbourne"}
	fmt.Println("City:   ", d.city)
	sold, name := calculateenvUtil(149, 200, d)
	fmt.Println("Dealer: ", name, "\nSold", sold)

}
