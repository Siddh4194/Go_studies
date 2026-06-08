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

func worker(ctx context.Context, id int, jobs <-chan int, appStats *AppStats, wg *sync.WaitGroup) {
	defer wg.Done()
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
	signalCtx, cancel := signal.NotifyContext(context.Background(),os.Interrupt,syscall.SIGTERM)
	defer cancel() // Good practice: always clean up resources

	ctx, timeoutCancel := context.WithTimeout(signalCtx,time.Second * 6)
	defer timeoutCancel()

	jobs := make(chan int, 10)

	wg := &sync.WaitGroup{}

	numWorkers := 2
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, appStats, wg)
	}


	for i := 1; i <= 5; i++ {
		jobs <- i
		appStats.IncrementTotal()
	}
	close(jobs)


	go func() {
		<-ctx.Done()
		fmt.Println("[Main] Shutdown signal received, waiting for active worker to wrap up.")
	}()
	wg.Wait()
	fmt.Printf("Main: App exited.\n Stats - Total: %d, Completed: %d, Failed: %d\n", appStats.TotalJobs, appStats.CompletedJobs, appStats.FailedJobs)
	
	time.Sleep(500 * time.Millisecond)
}
