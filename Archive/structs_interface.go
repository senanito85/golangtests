// Also available at: https://play.golang.org/p/yTTpB5gB6C

package main

import (
	"fmt"
)

// *****************************************************************************
// Example 1 - Struct vs Struct with Embedded Type
// *****************************************************************************

// Struct types contain fields. You can create instances of struct types.
// Struct with 2 fields
type Person struct {
	Name string // field
	Age  int    // field
}

// Struct with 2 fields
type Boy1 struct {
	Person      Person // field
	FavoriteToy string // field
}

// Struct with 2 fields (1 is an embeded type)
type Boy2 struct {
	Person             // anonymous field (embedded type)
	FavoriteToy string // field
}

// Struct with 3 fields (1 is an embeded type)
// Since Boy3 has an Age field, Person.Age is not longer a promoted field
type Boy3 struct {
	Person             // anonymous field (embedded type)
	FavoriteToy string // field
	Age         int    // field (named the same as Person.Age)
}

func example1() {
	// Create an instance of the Boy
	b1 := Boy1{Person{"Jon", 15}, "Corvette"}
	b2 := Boy2{Person{"Jon", 15}, "Corvette"}
	b3 := Boy3{Person{"Jon", 15}, "Corvette", 18}

	// Can ONLY access Age field 1 way
	fmt.Println(b1.Person.Age)
	// Output: 15

	// Can access Age field 2 ways (b2.Age is called a promoted field)
	fmt.Println(b2.Person.Age, b2.Age)
	// Output: 15 15

	// Can access Age field 2 ways (b3.Age is NOT a promoted field)
	fmt.Println(b3.Person.Age, b3.Age)
	// Output: 15 18

	// Can access FavoriteToy field 1 way in both
	fmt.Println(b1.FavoriteToy)
	// Output: Corvette
	fmt.Println(b2.FavoriteToy)
	// Output: Corvette
}

// *****************************************************************************
// Example 2 - Named Type Method Sets
// *****************************************************************************

// Create named Person type called Man
type Man Person

// Create named Person type called Woman
type Woman Person

// Create method increases the age
// Can ONLY be used with Man type
func (m *Man) IncreaseAge() {
	m.Age++
}

// Create method for only Person that decreases the age
// Can ONLY be used with Person type (not with Man or Woman)
// *** Can be used with Boy1 or Boy2
func (p *Person) DecreaseAge() {
	p.Age--
}

func example2() {
	// Create an instance of a Man
	r := Man{"Steve", 35}

	// This method can ONLY be called on a Man. It cannot be called on a Woman
	// or a Person.
	r.IncreaseAge()

	// Since Man is NOT a Person, this will throw a compile-time error
	//r.DecreaseAge()

	// Print 2 fields
	fmt.Println(r.Name, r.Age)
	// Output: Steve 36
}

// *****************************************************************************
// Example 3 - Usage of Interfaces
// *****************************************************************************

// Interface with a method set
type Animal interface {
	Speak() string
}

// Named interface is 100% compatible with Animal (not useful though, just good
// for aliasing third party interfaces)
type LandAnimal Animal

// Empty struct
type Dog struct {
}

// Empty struct
type Cat struct {
}

// Create method for the Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Receives any type that satisfies the Animal interface
func LeaveAtAnimalShelter(a Animal) {
	fmt.Printf("%T is now at the animal shelter.\n", a)
}

// This is invalid because a method cannot receive an interface
/*func (a Animal) Leave() {
	fmt.Printf("%T is now at the animal shelter.\n", a)
}*/

func example3() {
	// Create an instance of a Dog
	d := Dog{}

	// Call the method
	fmt.Println(d.Speak())
	// Output: Woof!

	// Call the function and pass the variable
	LeaveAtAnimalShelter(d)
	// Output: main.Dog is now at the animal shelter.

	// *************************************************************************
	// This is where it gets interesting
	// Can also create an instance of either interface by passing in the Dog
	a := Animal(d)
	l := LandAnimal(d)

	// Can even pass in an interface created from an interface that was passed
	// in the Dog
	p := LandAnimal(a)

	// Can call the method
	fmt.Println(p.Speak())
	// Output: Woof!

	// Can pass any of the instances that are compatible
	LeaveAtAnimalShelter(a)
	// Output: main.Dog is now at the animal shelter.
	LeaveAtAnimalShelter(l)
	// Output: main.Dog is now at the animal shelter.
	LeaveAtAnimalShelter(d)
	// Output: main.Dog is now at the animal shelter.
	LeaveAtAnimalShelter(p)
	// Output: main.Dog is now at the animal shelter.

	/*
		// Create an instance of a Cat
		c := Cat{}

		// This will throw a compile-time error because the Cat does NOT satisfy
		// the requirements to be an Animal.
		LeaveAtAnimalShelter(c)

		// Simple: In order for the Cat to be an Animal, it must have a Speak()
		// method.
		// Explanation: In order for the Cat to implement the Animal interface,
		// it must be associated, at a minimum, with the same method set as the
		// interface.
	*/
}

// *****************************************************************************
// Main
// *****************************************************************************

func main() {
	example1()
	example2()
	example3()
}
