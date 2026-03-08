package advanced

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// CONSTRUCTION SITE EXAMPLE
type worker struct {
	id int
}

// performTask simulates a worker performing task
func (w *worker) performTask(wg *sync.WaitGroup, tasks <-chan string) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker #%d started %s\n", w.id, task)
		time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
		fmt.Printf("Worker #%d finished %s\n", w.id, task)
	}
	fmt.Printf("Worker #%d finished\n", w.id)
}

func main() {

	var wg sync.WaitGroup

	// Define tasks to be performed by workers
	allTasks := []string{"digging", "laying bricks", "painting", "jumping", "sitting", "singing", "writing", "paying"}
	tasks := make(chan string, len(allTasks))
	for _, task := range allTasks {
		tasks <- task
	}
	close(tasks)

	// Define workers
	numWorkers := 3
	for i := range numWorkers {
		worker := worker{id: i + 1}
		wg.Add(1)
		go worker.performTask(&wg, tasks)
	}

	wg.Wait()

	fmt.Println("Tasks finished")
}

// === EXAMPLE WITH CHANNELS
// func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		fmt.Printf("Worker #%d started job #%d\n", id, job)
// 		time.Sleep(time.Second * time.Duration(id)) // simulate work
// 		results <- job * 2
// 		fmt.Printf("Worker #%d finished job #%d\n", id, job)
// 	}
// 	fmt.Printf("Worker #%d finished\n", id)
// }

// func main() {

// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, jobs, results, &wg)
// 	}

// 	for i := range numJobs {
// 		jobs <- i + 1
// 	}
// 	close(jobs)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}
// }

// === BASIC EXAMPLE WITHOUT CHANNELS
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// wg.Add(1) // WRONG PRACTICE
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // simulate work
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {

// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers) // CORRECT PRACTICE

// 	// Launch workers
// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait() // blocks until 3 calls of wg.Done()

// 	fmt.Println("All workers finished")
// }
