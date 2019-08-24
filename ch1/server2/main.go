// minimal echo and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8002", nil)) // nil is default handler
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // ensure only 1 go routine can access this at a time
	count++
	mu.Unlock()
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // ensure only 1 go routine can access this at a time
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
