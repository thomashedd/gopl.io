package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		//fmt.Println(i)
		//fmt.Println(arg)
	}
	//fmt.Println(os.Args[0])
	fmt.Println(s)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
