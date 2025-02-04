package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}

// Continue at https://quii.gitbook.io/learn-go-with-tests/build-an-application/json#write-the-test-first-1
