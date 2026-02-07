package main

import (
	"fmt"
	"sync"
	"time"
)

// ANSI color constants
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
)

func main() {
	fmt.Printf("%s==========================================%s\n", Cyan, Reset)
	fmt.Printf("%s🚀 Worker Pool Practice: Enhanced Output%s\n", Cyan, Reset)
	fmt.Printf("%s==========================================%s\n", Cyan, Reset)

	var wg sync.WaitGroup
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 workers
	workerColors := []string{Red, Green, Yellow}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, workerColors[i-1], jobs, results, &wg)
	}

	// Send 10 jobs
	fmt.Printf("%s[SYSTEM] Sending 10 jobs to the pool...%s\n", Gray, Reset)
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for completion
	wg.Wait()
	close(results)

	// Final Summary
	fmt.Printf("\n%s==========================================%s\n", Purple, Reset)
	fmt.Printf("%s📊 EXECUTION SUMMARY%s\n", Purple, Reset)
	fmt.Printf("%s==========================================%s\n", Purple, Reset)
	for ans := range results {
		fmt.Printf("✅ %sResult: %-3d%s\n", Gray, ans, Reset)
	}
	fmt.Printf("%s==========================================%s\n", Purple, Reset)
}

func worker(id int, color string, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		timestamp := time.Now().Format("15:04:05")
		fmt.Printf("[%s] %sWorker %d%s | 🛠️  Processing job %-2d\n", timestamp, color, id, Reset, j)

		time.Sleep(time.Duration(500+(id*100)) * time.Millisecond) // Simulate slightly different processing times

		fmt.Printf("[%s] %sWorker %d%s | ✅ Finished job %-2d\n", timestamp, color, id, Reset, j)
		results <- j * 2
	}
}
