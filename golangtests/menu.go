package main

import (
	"fmt"
	lib "ifelse"
	libtwo "psclies"
)

func main() {
	Starter()
}

func Starter() {
	multiline2 := "Welcome! Please make your choice: \n" + "------------------------------ \n" + "1. Option One \n" + "2. Option Two \n" +
		"3. Option Three \n" + "0. Exit \n" + "----------------- \n" + "Your Choice is: "
	fmt.Print(multiline2)
	var usersChoice int
	fmt.Scanln(&usersChoice)
	decider(usersChoice)
}

func decider(value int) {
	switch value {
	case 1:
		fmt.Println("You Chose One, try again \n")
		lib.Tester()
		Starter()
	case 2:
		fmt.Println("You Chose Two, try again \n")
		libtwo.Setter()
		Starter()
	case 3:
		fmt.Println("You Chose Three, try again \n")
		Starter()
	case 0:
		fmt.Println("You Chose To Quit")
		fmt.Println("-----------------")
		fmt.Println("Good Bye")
	}
}
