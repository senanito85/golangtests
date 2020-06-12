package main

import (
	"fmt"
	"os"
	"github.com/gocarina/gocsv"
)

type Car struct {
	Year  string `csv:"YEAR"`
	Make  string `csv:"MAKE"`
	Model string `csv:"MODEL"`
	Km    string `csv:"KM"`
}


func main() {

    carsFile, err := os.Open("cars.csv")
    if err != nil {
        panic(err)
    }
    defer carsFile.Close()

    cars := []Car{}

    if err := gocsv.UnmarshalFile(carsFile, &cars); err != nil {
        panic(err)
    }
    /*
    for _, car := range cars { //prints everything
        fmt.Println(car)
    }
    */
    fmt.Println("------------------------------------")

    var carToLookFor string = "Holden"
    var found bool
    var number int = 0

    for _, car := range cars {
    	if car.Make == carToLookFor {
    		found = true;
    		fmt.Println(car)
    		number ++
    	}
        
    } 
    fmt.Printf("\nFound? %v - %d Cars found", found, number)
}

