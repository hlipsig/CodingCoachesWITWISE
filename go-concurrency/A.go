package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// A calls all the URLs in serial. While this works, it takes a long time to run.
func A(urlsFile string) {
	startTime := time.Now()

	urls := strings.Split(urlsFile, "\n")

	for _, url := range urls {
		resp, err := http.Get("https://" + url)
		if err != nil {
			fmt.Printf("error with %s: %v\n", url, err)
		} else {
			fmt.Printf("received code %d from %s\n", resp.StatusCode, url)
		}
	}

	fmt.Printf("Executed in %f seconds\n", time.Now().Sub(startTime).Seconds())
}
