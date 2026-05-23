package main

// import (
// 	"fmt"
// 	"time"
// )

// func printMessage(msg string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(msg)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	go printMessage("Hello")
// 	go printMessage("Go")

// 	// Wait long enough for goroutines to finish
// 	time.Sleep(6 * time.Second)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func printMessage(msg string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(msg)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go printMessage("Hello", &wg)
// 	go printMessage("Go", &wg)

// 	wg.Wait()
// }

import (
	"fmt"
	"time"
)

func worker (id int,jobs chan int, results chan int){
    for j := range jobs {
        fmt.Printf("Worker %d started the job %d\n",id,j);
        time.Sleep(time.Second)
        fmt.Printf("Worker %d finished the job %d\n",id,j);
        results <- j * j
    }
}

// var urls = []string{
//     "https://google.com",
//     "https://github.com",
//     "https://httpbin.org/status/200",
//     "https://httpbin.org/status/404",
//     "https://httpbin.org/status/500",
//     "https://httpbin.org/delay/5",
//     "https://invalid-domain-test-123.com",
// }

// func main() {
//     Intaker(urls)    
    // const numJobs = 5
    // const numWorkers = 3

    // jobs := make(chan int,numJobs)
    // results := make(chan int,numJobs)

    // for w := 1 ; w < numWorkers; w++ {
    //     go worker(w,jobs,results)
    // }

    // for j := 1; j < numJobs; j++ {
    //     jobs <- 1
    // }
    // close(jobs)

    // for r := 1; r < numJobs; r++{
    //     fmt.Println("Result received ",<-results)
    // }
// }
