package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// B introduces goroutines to our loop, and it runs really fast!
// Goroutines are lightweight threads managed by the Go runtime.
// For more info on Goroutines, see https://go.dev/tour/concurrency/1
// ... but we don't see any results. Can you spot what's wrong?
func B(urlsFile string) {
	startTime := time.Now()

	urls := strings.Split(urlsFile, "\n")

	for _, url := range urls {
		// The syntax `go` represents the creation of a goroutine.
		// We then create a new anonymous function with `func(){...}()` to be executed by the goroutine.
		go func() {
			resp, err := http.Get("https://" + url)
			if err != nil {
				fmt.Printf("error with %s: %v\n", url, err)
			} else {
				fmt.Printf("received code %d from %s\n", resp.StatusCode, url)
			}
		}()
	}

	fmt.Printf("Executed in %f seconds\n", time.Now().Sub(startTime).Seconds())
}
