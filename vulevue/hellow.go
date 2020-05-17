package main

import "fmt"

func main() {
	y := 11
	var ar [10]int
	brr := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := 0; i < len(ar); i++ {
		ar[i] = y
		y = 1 + y
	}
	fmt.Println(ar)
	fmt.Println(brr)

	//----------------------------------------------------------------
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\n--------------- Example 1 --------------------\n")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	fmt.Println("\n---------------Example 3--------------------\n")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\n---------------Example 4--------------------\n")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}
