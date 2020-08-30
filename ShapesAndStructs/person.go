package main

import "fmt"

type Employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

type Person interface {
	Introduce() string
}

func (e Employee) Introduce() string {
	return e.firstName
}

func main() {
	ross := &Employee{
		firstName: "Ross",
		lastName:  "Bing",
		salary:    1200,
		fullTime:  true,
	}

	fmt.Println("firstName", (*ross).firstName)
	fmt.Println("lastName", (*ross).lastName)

	var a Person = Employee{"Sanan", "Ibrahimoff", 5500, true}
	fmt.Println(a.Introduce())
	fmt.Println(a)

}
