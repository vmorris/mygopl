// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	// how many digits before the first comma?
	prefix := len(s) % 3
	if prefix == 0 {
		prefix = 3
	}
	buf.WriteString(s[:prefix])

	// now insert commas
	for i := prefix; i < len(s); i += 3 {
		buf.WriteString("," + s[i:i+3])
	}
	return buf.String()
}
