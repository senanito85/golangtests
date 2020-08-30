package main

import (
	"fmt"

	L "ifelse"
	M "psclies"
)

func main() {
	starter()
}

func starter() {
	multiline2 := "Welcome! Please make your choice: \n" +
		"--------------------------------- \n" +
		"1. Option One \n" + "2. Option Two \n" +
		"3. Option Three \n" + "0. Exit \n" +
		"----------------- \n" + "Your Choice is: "
	fmt.Print(multiline2)
	var usersChoice int
	fmt.Scanln(&usersChoice)
	decider(usersChoice)
}

func decider(value int) {
	switch value {
	case 1:
		fmt.Println("You Chose One, try again")
		L.Tester()
		starter()
	case 2:
		fmt.Println("You Chose Two, try again")
		M.Setter()
		starter()
	case 3:
		fmt.Println("You Chose Three, try again")
		starter()
	case 0:
		fmt.Println("You Chose To Quit")
		fmt.Println("Good Bye")
	}
}
