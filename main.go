package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Worker Pool Practice")

	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Create channels for jobs and results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		// Increment the counter for each goroutine
		wg.Add(1)

		// Start a worker goroutine
		go worker(i, jobs, results, &wg)
	}

	// Send 10 jobs to the jobs channel
	for j := 1; j <= 10; j++ {
		// Send a job to the jobs channel
		jobs <- j
	}
	// Close the jobs channel to signal that no more jobs will be sent
	close(jobs)

	// Wait for all worker goroutines to complete
	wg.Wait()
	// Close the results channel to signal that no more results will be sent
	close(results)

	// Print all the results
	for ans := range results {
		fmt.Println("Result:", ans)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	// Decrement the counter when the goroutine completes
	defer wg.Done()

	// Loop over the jobs channel and process each job
	for j := range jobs {
		// Print the job being processed
		fmt.Println("Worker", id, "processing job", j)

		// Send the result to the results channel
		results <- j * 2
	}
}
