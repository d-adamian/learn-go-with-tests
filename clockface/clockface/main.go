package main

import (
	"learn-go-with-tests/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}

// Currently at https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/math#refactor-5
