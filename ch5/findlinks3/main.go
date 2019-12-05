package main

import (
	"fmt"
	"log"
	"os"

	links "github.com/thomashedd/gopl.io/ch5/links"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadFirst(crawl, os.Args[1:])
}

//breadFirst calls f for each item in the worklist.
//Any items returned by f are added to the worklist.
// f is called at most once for each item
func breadFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...) // "f(item)..." causes all the items in the list returned by f to be appened to the worklist
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
