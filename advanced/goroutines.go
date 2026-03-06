package advanced

import (
	"fmt"
	"time"
)

func main() {

	var err error

	fmt.Println("Beginning program")
	go sayHello()
	fmt.Println("After sayHello function")

	go printNumbers()
	go printLetters()
	// err = go doWork()
	go func() {
		err = doWork()
	}()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := range 5 {
		fmt.Printf("%-15s %v\n", fmt.Sprintf("Number: %d", i), time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Printf("%-15s %v\n", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("An error occurred in doWork")
}
