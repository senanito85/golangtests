package main
import ( "fmt" )

func countAverage(numbers ...int) int {
	sum := 0
	for _, i := range numbers {
		sum = sum + i
	}
	return sum / len(numbers)
}

func countSum(numbers ...int) int {
	sum := 0
	for _, i := range numbers {
		sum = sum + i
	}
	return sum
}

func findSmallest(numbers ...int) int {
	smallest := numbers[0]
	for _, i := range numbers {
		if i < smallest {
		smallest = i	
		}
	}
    return smallest
}

func findBiggest(numbers ...int) int {
	biggest := numbers[0]
	for _, i := range numbers {
		if i > biggest {
		biggest = i	
		}
	}
    return biggest
}

func main() {

        //---------------------------------------------

        average := countAverage(13, 10, 14, 17, 12, 16)
        summary := countSum(13, 10, 14, 17, 12, 16)
        smallest := findSmallest(13, 10, 14, 17, 12, 16)
        biggest := findBiggest(13, 10, 14, 17, 12, 16)
        
        fmt.Println("Average of numbers is ", average)
        fmt.Println("Sum of numbers is ", summary)
        fmt.Println("Smallest number is ", smallest)
        fmt.Println("Biggest number is ", biggest)

}

 