// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"fmt"
	"os"
)

func main() {
	cl := os.Args[1:]
	//s := strings.Join(cl, " ")
	s := string(cl)
	fmt.Println(comma(s))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
