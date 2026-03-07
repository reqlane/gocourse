package advanced

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// Simulate processing of ticket requests
func ticketProcessor(id int, requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("%d.%d Worker %d processing person ID %d with %d ticket(s) with total cost %d\n", time.Now().Second(), time.Now().Nanosecond(), id+1, req.personID, req.numTickets, req.cost)
		// Simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func main() {

	price := 5

	numWorkers := 3
	numRequests := 10
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int, numRequests)

	// Start ticket processor/worker
	for i := range numWorkers {
		go ticketProcessor(i, ticketRequests, ticketResults)
	}

	// Send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	// Process results
	for range numRequests {
		result := <-ticketResults
		fmt.Printf("%d.%d Ticket for person ID %d processed successfully!\n", time.Now().Second(), time.Now().Nanosecond(), result)
	}
}

// === BASIC WORKER POOL PATTERN
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d proccessing task %d\n", id, task)
// 		// Simulate work
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// func main() {

// 	numWorkers := 3
// 	numTasks := 10

// 	tasks := make(chan int, numTasks)
// 	results := make(chan int, numTasks)

// 	// Create workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	// Send values to the tasks channel
// 	for i := range numTasks {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	// Collect the results
// 	for range numTasks {
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}
// }
