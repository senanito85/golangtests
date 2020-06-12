package main
import ( "fmt" )

type individ interface {
	intro() string
}

type person struct {
	name string
	surname string
	age int
	location string
}

type employee struct {
	employeId int
	department string
	slary int
	human person
}

func (e employee) intro() string {
	return e.department
}

func (p person) intro() string {
	return p.name
}

func speakSmthn(id individ) string {
	return id.intro()
}

func main() {
	p1 := person{"Sanan", "Ibra", 34, "Perth"}
	p2 := person{"Jamie", "Pan", 26, "Perth"}
	e1 := employee{1007, "IT", 155000, p1}
	e2 := employee{1008, "Fiance", 115000, p2}

	team := []individ{e1, e2}
    crew := []individ{p1, p2}

    for _, i := range crew {
		fmt.Println("My name is", speakSmthn(i))
	}

	for _, i := range team {
		fmt.Println("I work in", i.intro())	 //fmt.Println(i)	
	}
}