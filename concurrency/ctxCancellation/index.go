package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type AppStats struct {
	mu 		sync.Mutex
	TotalJobs int
	CompletedJobs int
	FailedJobs int
}

func (app *AppStats) IncrementTotal() {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.TotalJobs++
}

func (app *AppStats) IncrementCompleted() {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.CompletedJobs++
}

func (app *AppStats) IncrementFailed() {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.FailedJobs++
}

func worker(ctx context.Context, id int, jobs <-chan int, appStats *AppStats) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Stopping due to: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// Simulate a task
			fmt.Printf("Worker %d: Starting job %d\n", id, job)
			
			// We simulate work but keep an eye on the context
			processJob(ctx, job, appStats)
		}
	}
}

func processJob(ctx context.Context, job int, appStats *AppStats) {
	select {
	case <-time.After(2 * time.Second): // Simulate 2s work
		fmt.Printf("Finished job %d\n", job)
		appStats.IncrementCompleted()
	case <-ctx.Done():
		// 2. TIMEOUT PROPAGATION: Work stopped early
		fmt.Printf("Job %d aborted\n", job)
		appStats.IncrementFailed()
	}
}

func main() {

	// init app stats
	appStats := &AppStats{}

	// 3. SET TIMEOUT: The whole app has 3 seconds to work
	ctx, cancel := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer cancel() // Good practice: always clean up resources

	jobs := make(chan int, 10)
	for i := 1; i <= 5; i++ {
		jobs <- i
		appStats.IncrementTotal()
	}

	go worker(ctx, 1, jobs, appStats)

	// Wait to see the timeout happen
	<-ctx.Done()

	

	time.Sleep(500 * time.Millisecond) // Brief pause to see logs
	fmt.Println("Main: App exited.")
}
