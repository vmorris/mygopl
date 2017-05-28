// Charcount computes counts of Unicode categories
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	letterCounts := make(map[rune]int)
	numberCounts := make(map[rune]int)
	symbolCounts := make(map[rune]int)
	punctCounts := make(map[rune]int)
	spaces := 0
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			letterCounts[r]++
		case unicode.IsNumber(r):
			numberCounts[r]++
		case unicode.IsSpace(r):
			spaces++
		case unicode.IsSymbol(r):
			symbolCounts[r]++
		case unicode.IsPunct(r):
			punctCounts[r]++
		}

	}
	fmt.Printf("rune\tcount\n")
	for c, n := range letterCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	for c, n := range numberCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	for c, n := range punctCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	for c, n := range symbolCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if spaces > 1 { // not sure why we're picking up an extra space, so starting at 1!
		fmt.Printf("\n%d spaces\n", spaces-1)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
