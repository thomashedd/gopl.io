// boiling prints the boiling point of water
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gdegreesF or %gdegreesC\n", f, c)
	//output:
	// boiling point = 212 degreesF or 100 degrees C
}
