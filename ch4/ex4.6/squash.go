package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "this   is     a			 test"
	s = squash(s)
	fmt.Println(s)
}

func squash(str string) string {
	bytes := []byte(str)
	//log.Println("input = " + string(bytes))
	out := bytes[:0] // empty slice of bytes
	spaceAdded := false
	for _, b := range bytes {
		//log.Println("checking " + string(b))
		if unicode.IsSpace(rune(b)) {
			//log.Println("found a space")
			if !spaceAdded {
				//log.Println("appending a space")
				out = append(out, ' ')
				spaceAdded = true
			}
		} else {
			out = append(out, b)
			spaceAdded = false
		}
	}
	return string(out)
}
