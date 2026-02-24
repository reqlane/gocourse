package basics

import "fmt"

func main() {

	// Simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// continue, break
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Number:", i)
		if i == 5 {
			break
		}
	}

	// Star pyramid
	rows := 5

	for i := 1; i <= rows; i++ {
		// Inner loop for spaces before stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// Inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Range
	for i := range 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println("We have a lift off!")
}
