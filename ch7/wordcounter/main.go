package main

import (
	"fmt"
	"bufio"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	*c += WordCounter()
}