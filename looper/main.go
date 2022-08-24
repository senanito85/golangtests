package main

import "fmt"

var candidates = []string{"Sanan", "Jamie", "Tiffany", "Ralph"}

func main() {

	fmt.Println("\n---------------Example 0 --------------------\n")
	for _, value := range candidates {
		fmt.Println("Hi ", value)
	}

	fmt.Println("\n---------------Example 1 --------------------\n")
	for i, e := range candidates {
		fmt.Println(i, "--", e)
	}

	j := 0
	fmt.Println("\n---------------Example 2 --------------------\n")
	for range candidates {
		fmt.Println(candidates[j])
		j++
	}
}
