// Wordfreq reports the frequency of each word in an input text file.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	counts := make(map[string]int)

	file, err := os.Open("sample3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// for each word assign it as a key in the map
	// for each word increment the value in the map
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	//fmt.Printf("word frequency:\n")
	//for word, occurrence := range counts {
	//	fmt.Printf("%s : %d\n", word, occurrence)
	//}

	words := make([]string, 0, len(counts))

	for word := range counts {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool { return counts[words[i]] < counts[words[j]] })

	for _, word := range words {
		fmt.Println(word)
	}
	//fmt.Println(words)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
