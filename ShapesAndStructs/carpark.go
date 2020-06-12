package main

import "fmt"

type Car struct {
	year  int
	make  string
	model string
	km    int
}

func (c *Car) String() string {
	return fmt.Sprintf("\nMake: %s\nModel: %s\nYear: %d\nKm's: %d\n", c.make, c.model, c.year, c.km)
}

func main() {

	myCar := Car{year: 1996, make: "Toyota", model: "LandCrouiser", km: 260000}
	scCar := Car{year: 1998, make: "Mersedes", model: "CLK 320", km: 230000}
	drCar := Car{year: 2001, make: "Chevrolet", model: "Malibu", km: 190000}
	bxCar := Car{year: 2004, make: "BMW", model: "X5", km: 165000}
	auCar := Car{year: 2009, make: "Holden", model: "Trax", km: 15000}

	cars := [4]Car{myCar, scCar, drCar, bxCar}

	var soldCars []Car = cars[0:4] //creates a slice from a[0] to a[3]

	soldCars = append(soldCars, auCar)

	for _, car := range soldCars {
		//fmt.Printf("\nMake: %s\nModel: %s\nYear: %d\nKm's: %d\n", car.make, car.model, car.year, car.km)
		fmt.Println(car)
	}

	for i := 0; i < len(soldCars); i++ {
		fmt.Println(soldCars[i])
	}

}
