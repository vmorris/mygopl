// diffsha256 counts the number of bits that are different between
// two sha256 hashes created from CLI args
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// pc[i] is the population count of i.
var pc [32]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	usage := "usage: ./diffsha256 string1 string2"
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))

	sum := 0

	for i, b1 := range c1 {
		b2 := c2[i]
		x := b1 ^ b2
		for x != 0 { // if not zero, clear and add 1
			x = x & (x - 1)
			sum++
		}
	}

	fmt.Printf("The number of bits that differ between the two hashes is: %d\n", sum)

}
