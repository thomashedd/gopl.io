package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	m := make(map[string]int)

	for k, v := range check(m, doc) {
		fmt.Printf("%s:%d\n", k, v)
	}
}

func check(m map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return m
	} else {
		if n.Type == html.ElementNode {
			m[n.Data]++
		}
		m = check(m, n.FirstChild)
		m = check(m, n.NextSibling)
	}
	return m
}
