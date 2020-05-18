package main

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s\n", p.name)
	fmt.Printf("and my age is %d\n", p.age)
	fmt.Printf("----------------\n")
}

func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	//p := person{name: "Sanan", age: 34, alive: true}
	//fmt.Println(p)
	//
	NewPerson("Sanan", 35).Greet()
	NewPerson("Jame", 26).Greet()
}
