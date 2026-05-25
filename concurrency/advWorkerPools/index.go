package main

import (
	"fmt"
	"sync"
)

// fixed worker pools
func main(){
	jobs := make (chan int,100)
	workers := 5

	result := make (chan string,100)
	var wg sync.WaitGroup
	for w := 1; w < workers; w++ {
		wg.Add(1)
		go func(w int)  {
			defer wg.Done()
			for j := range jobs {
				result <-  fmt.Sprintf("Worker %d started the job %d\n",w,j)			}
		}(w)
	}

	for j := 1; j < 100; j++ {
		jobs <- j
	}

	//  by using this we tell the goroutines that there are no more jobs to process
	close(jobs)
	wg.Wait()

	// if we don't close the result channel then the main goroutine will be waiting for the results to come in and it will never exit
	close(result)

	for r := 1; r < 100; r++ {
		fmt.Println(<-result)
	}
}