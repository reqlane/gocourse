package basics

import "fmt"

func main() {

	message := "Hello World"

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		slice[i] = 99
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	fmt.Println(slice)
}
