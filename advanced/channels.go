package advanced

import (
	"fmt"
	"time"
)

func main() {

	// varName := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // blocking because continuously trying to receive values from a stream
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()
	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		receiver = <-greeting
		fmt.Println(receiver)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program")
}
