package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// D replaces our `time.Sleep` call with a WaitGroup.
// WaitGroups in Go make a lot of sense to read, but you can find more information here: https://gobyexample.com/waitgroups
func D(urlsFile string) {
	startTime := time.Now()

	urls := strings.Split(urlsFile, "\n")

	wg := sync.WaitGroup{}
	for _, url := range urls {
		// This line tells the wait group that we have 1 additional thread executing.
		wg.Add(1)
		go func() {
			// This line will tell the wait group that the thread's execution is complete once it finishes.
			defer wg.Done()
			resp, err := http.Get("https://" + url)
			if err != nil {
				fmt.Printf("error with %s: %v\n", url, err)
			} else {
				fmt.Printf("received code %d from %s\n", resp.StatusCode, url)
			}
		}()
	}

	// This line forces our main thread to wait here until all threads have completed their execution.
	wg.Wait()

	fmt.Printf("Executed in %f seconds\n", time.Now().Sub(startTime).Seconds())
}
