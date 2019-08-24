package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
