package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		w, i, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Println("CountWordsAndImages error: ", err)
		}
		fmt.Printf("words = %d,images = %d\n", w, i)
	}
	//fmt.Printf("words = %d,images = %d\n", w, i)
}

//CountWordsAndImages does an HTTP GET request for the HMTL
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsandImages(doc)
	return
}

func countWordsandImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	} else {
		if n.Type == html.ElementNode {
			if n.Data == "img" {
				images++
			}
		} else if n.Type == html.TextNode {
			scanner := bufio.NewScanner(strings.NewReader(n.Data))
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				words++
			}
		}
		w1, i1 := countWordsandImages(n.FirstChild)
		words += w1
		images += i1
		w2, i2 := countWordsandImages(n.NextSibling)
		words += w2
		images += i2
	}
	return
}
