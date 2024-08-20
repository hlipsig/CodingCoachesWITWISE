package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// C adds a call to `time.Sleep` to make sure our main thread waits for all of our threads to be finished.
// This has a number of problems - we're either waiting too long or potentially not long enough.
// Note here that our results are all printed out long before the program prints its final line and exits.
func C(urlsFile string) {
	startTime := time.Now()

	urls := strings.Split(urlsFile, "\n")

	for _, url := range urls {
		go func() {
			resp, err := http.Get("https://" + url)
			if err != nil {
				fmt.Printf("error with %s: %v\n", url, err)
			} else {
				fmt.Printf("received code %d from %s\n", resp.StatusCode, url)
			}
		}()
	}

	time.Sleep(10 * time.Second)

	fmt.Printf("Executed in %f seconds\n", time.Now().Sub(startTime).Seconds())
}
