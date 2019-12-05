// reverse reverses a slice of ints in place.
package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int) {
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	rotate(b[:])
	//fmt.Println(a) // "[5 4 3 2 1]""
	fmt.Println(b)

}
