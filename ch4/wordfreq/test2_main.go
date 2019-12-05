package main

import (
	"fmt"
)

func main() {
	file, err := os.Open(sample.txt)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}