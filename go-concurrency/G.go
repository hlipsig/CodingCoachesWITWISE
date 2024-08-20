package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// G while F will work totally fine for this program, we could run into issues in a more complex code base.
// For example, what if one of our threads has an unforeseen error and never writes a result to the channel?
// In such an instance, our printing for-loop would wait forever as it never would reach the number of expected results.
// This function combines both concepts we've learned about - WaitGroups and channels - to steal benefits from both.
// We use a WaitGroup to keep track of how many threads are running, and channels to handle thread safety.
// We also use an extra goroutine to process the WaitGroup, and manually close the channel when all of our goroutines have completed.
func G(urlsFile string) {
	startTime := time.Now()

	urls := strings.Split(urlsFile, "\n")

	wg := sync.WaitGroup{}
	resultsChan := make(chan string)
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get("https://" + url)
			if err != nil {
				resultsChan <- fmt.Sprintf("error with %s: %v", url, err)
			} else {
				resultsChan <- fmt.Sprintf("received code %d from %s", resp.StatusCode, url)
			}
		}()
	}

	// This block spins up a separate goroutine that will wait until all threads have finished execution and then call "close" on the channel.
	// Once the channel is closed, the for-loop below knows that there is no more data to be read.
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		fmt.Println(result)
	}

	fmt.Printf("Executed in %f seconds\n", time.Now().Sub(startTime).Seconds())
}
