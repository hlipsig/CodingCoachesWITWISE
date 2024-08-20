package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// E addresses the concern with writing our results in a thread safe manner.
// Writes are not guaranteed to be atomic, which means that two threads could start simultaneously writing to the terminal and have their results mixed.
// There are multiple ways we can fix this issue, but here we're going to explore channels.
// Channels are typed queues that you can pass data through in a thread-safe manner.
// You can find more info about channels here: https://go.dev/tour/concurrency/2
// Note that the program never exits, we must be missing something...
func E(urlsFile string) {
	startTime := time.Now()

	urls := strings.Split(urlsFile, "\n")

	// Create a new channel - you can specify buffer lengths for channels, but we'll ignore that for this exercise
	resultsChan := make(chan string)
	for _, url := range urls {
		go func() {
			resp, err := http.Get("https://" + url)
			// An arrow <- pointing at the channel variable is used to show that we are writing data to the channel
			if err != nil {
				resultsChan <- fmt.Sprintf("error with %s: %v", url, err)
			} else {
				resultsChan <- fmt.Sprintf("received code %d from %s", resp.StatusCode, url)
			}
		}()
	}

	// Go supports looping over channels - continually reading from the channel into the loop variable (`result` in this case)
	for result := range resultsChan {
		fmt.Println(result)
	}

	fmt.Printf("Executed in %f seconds\n", time.Now().Sub(startTime).Seconds())
}
