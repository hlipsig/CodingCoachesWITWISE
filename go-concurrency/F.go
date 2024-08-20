package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// F In E we were looping over the channel, but our program doesn't know when to stop looping!
// There are multiple ways we could solve this, but let's start with a naive solution.
// Since we know how many results we expect (one per URL), we can quit our printing loop after that many results.
func F(urlsFile string) {
	startTime := time.Now()

	urls := strings.Split(urlsFile, "\n")

	resultsChan := make(chan string)
	for _, url := range urls {
		go func() {
			resp, err := http.Get("https://" + url)
			if err != nil {
				resultsChan <- fmt.Sprintf("error with %s: %v", url, err)
			} else {
				resultsChan <- fmt.Sprintf("received code %d from %s", resp.StatusCode, url)
			}
		}()
	}

	for i := 0; i < len(urls); i++ {
		// Note that we need to explicitly read from the channel now with an arrow pointing out of the variable.
		fmt.Println(<-resultsChan)
	}

	fmt.Printf("Executed in %f seconds\n", time.Now().Sub(startTime).Seconds())
}
