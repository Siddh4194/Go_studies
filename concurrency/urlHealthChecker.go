package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	URL string
	Status bool
	Error string
	StatusCode int
	Latency time.Duration
}

func healthChecker(url string) Result {
	start := time.Now()
	
	client := http.Client{
		Timeout: 3* time.Second,
	}
	resp,err := client.Get(url)
	if err != nil {
		return Result{URL: url, Status: false, Error: err.Error(),StatusCode: 0, Latency: time.Since(start)}
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusOK {
		return Result{URL: url, Status: true, Error: "", StatusCode: resp.StatusCode, Latency: time.Since(start)}
	}
	return Result{URL: url, Status: false, Error: "Status code is not 200", StatusCode: resp.StatusCode, Latency: time.Since(start)}
}

func urlWorker(results chan Result, jobs chan string){
	for url := range jobs {
		results <- healthChecker(url)
	}
}

func Intaker(urls []string) {
	start := time.Now()

	defer func(){
		elapsed := time.Since(start)
		fmt.Printf("Time taken: %s \n",elapsed)
	}()

	numWorkers := 30
	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))

	for w := 1; w <= numWorkers; w++ {
		go urlWorker(results,jobs)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	for _,url := range urls {
		status := <- results
		fmt.Printf(
	"Status of %s is %t \n Error: %s \n Status Code: %d \n Latency: %s \n\n",
	url,
	status.Status,
	status.Error,
	status.StatusCode,
	status.Latency,
)}
}


var urls = []string{
    "https://google.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
    "https://github.com",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://httpbin.org/delay/5",
    "https://invalid-domain-test-123.com",
}

func main() {
    Intaker(urls)    }