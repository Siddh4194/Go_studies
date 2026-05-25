package main

import (
	"fmt"
	"sync"
)

// fixed worker pools
// func main(){
// 	jobs := make (chan int,100)
// 	workers := 5

// 	result := make (chan string,100)
// 	var wg sync.WaitGroup
// 	for w := 1; w < workers; w++ {
// 		wg.Add(1)
// 		go func(w int)  {
// 			defer wg.Done()
// 			for j := range jobs {
// 				result <-  fmt.Sprintf("Worker %d started the job %d\n",w,j)			}
// 		}(w)
// 	}

// 	for j := 1; j < 100; j++ {
// 		jobs <- j
// 	}

// 	//  by using this we tell the goroutines that there are no more jobs to process
// 	close(jobs)
// 	wg.Wait()

// 	// if we don't close the result channel then the main goroutine will be waiting for the results to come in and it will never exit
// 	close(result)

// 	for r := 1; r < 100; r++ {
// 		fmt.Println(<-result)
// 	}
// }

// BackPressure Worker Pools
// func uploadPools(job chan int, id int){
// 	select {
// 	case job <- id:
// 		fmt.Printf("Worker %d started the job %d\n",id,id);
// 	default:
// 		fmt.Printf("Queue is full message rejected for worker");
// 	}
// }
// func main(){
// 	job := make(chan int,5)

// 	uploadPools(job,1)
// 	uploadPools(job,1)
// 	uploadPools(job,1)
// 	uploadPools(job,1)
// 	uploadPools(job,1)
// 	uploadPools(job,1)
// }

// dynamic worker pools

func main(){
	jobs := make(chan int)
	maxWorkers := 10

	currentWorkers := 0

	var mu sync.Mutex

	for j := 1; j < 100; j++ {
		jobId := j

		select {
			case jobs <- jobId:
				fmt.Printf("Job %d accepted by a worker\n", jobId)
			default:
				mu.Lock()
				if currentWorkers < maxWorkers {
					currentWorkers++

					go func (workerId int) {
						fmt.Printf("Started new dynamic worker")
						for tasks := range jobs {
							fmt.Printf("Worker %d processing job %d\n", workerId, tasks)
						}
					}(currentWorkers)
					jobs <- jobId
					fmt.Printf("Job %d accepted by a new worker\n", jobId)
				} else {
					fmt.Println("Max workers reached, blocking")
					jobs <- jobId
				}
				mu.Unlock()
		}
	}
}