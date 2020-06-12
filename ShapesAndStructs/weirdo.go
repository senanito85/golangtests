package main

import "fmt"

type User struct {
	Name string `example:"name"`
}

type Car struct {
	year  int
	make  string
	model string
}

func (u *User) myString() string {
	return fmt.Sprintf("Hi! My name is %s", u.Name)
}

func (p *User) anotherString() string {
	return fmt.Sprintf("Hi! My name is not really %s", p.Name)
}

func (c *Car) String() string {
	return fmt.Sprintf("Make: %s, Model: %s, Year: %d", c.make, c.model, c.year)
}

func main() {
	u := &User{
		Name: "Sammy",
	}

	fmt.Println(u.myString())

	p := &User{
		Name: "Hammer",
	}

	fmt.Println(p.anotherString())

	myCar := &Car{year: 1996, make: "Toyota", model: "LandCrouiser"}
	fmt.Println(myCar)

	helloWorld := "helloworld"
	var pointerToHelloWorld *string
	pointerToHelloWorld = &helloWorld
	fmt.Println("Pointer To helloWorld")
	fmt.Println(*pointerToHelloWorld)

}
