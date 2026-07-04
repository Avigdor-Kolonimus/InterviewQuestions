package main

// Write code that prints HTTP response status codes
// for requests to two URLs:
// Google homepage and Wildberries homepage.
// Requests must be executed in parallel.

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetch(ctx context.Context, url string) int {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func withoutWorkerPool(urls []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	results := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			code := fetch(ctx, u)
			results <- fmt.Sprintf("%s -> %d", u, code)
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}

// Worker Pool version
func worker(ctx context.Context, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			results <- fmt.Sprintf("%s -> error", url)
			continue
		}

		results <- fmt.Sprintf("%s -> %d", url, resp.StatusCode)
		resp.Body.Close()
	}
}

func withWorkerPool(urls []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	jobs := make(chan string)
	results := make(chan string, len(urls))

	var wg sync.WaitGroup

	// number of workers = concurrency limit
	const workers = 5

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ctx, jobs, results, &wg)
	}

	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.wildberries.ru",
	}

	withoutWorkerPool(urls)
	withWorkerPool(urls)
}
