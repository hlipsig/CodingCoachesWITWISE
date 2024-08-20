package main

import (
	_ "embed"
)

//go:embed urls
var urlsFile string

func main() {
	A(urlsFile)
}
