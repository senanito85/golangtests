package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Person struct {
	fname string
	lname string
	age   int
}

func (p Person) hello() string {
	return "Hello, my name is " + p.fname + " " + p.lname + " Im " + strconv.Itoa(p.age)
}

func main() {
	//Instanciating users
	p1 := Person{"Sanan", "Ibrahimoff", 34}
	p2 := Person{"Xiaofen", "Pan", 26}
	p3 := Person{"Di", "Davali", 13}
	p4 := Person{"Saha", "Gray", 28}
	p5 := Person{"Masha", "Medvedova", 6}

	classroom := []Person{p1, p2, p3}
	classroom = append(classroom, p4, p5)

	if fileExists("db.txt") {
		//inform the user that file exists and start writing to file
		fmt.Println("DB file exists")

		f, err := os.OpenFile("db.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		//iterate through the array an write its content to a file.
		for i := 0; i < len(classroom); i++ {
			//fmt.Println(classroom[i].hello())
			if _, err := f.WriteString(classroom[i].hello() + "\n"); err != nil {
				log.Println(err)
			}
		}

	} else {
		//if file does not exist throw and error and don't do anything
		fmt.Println("DB file does not exist (or is a directory)")
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
