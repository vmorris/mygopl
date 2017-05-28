// reverse the characters of a []byte slice in place without allocating
// additional memory
package main

import "fmt"

func main() {
	in := "racecar's are neat!"
	inbytes := []byte(in)
	fmt.Println(string(inbytes))
	inbytes = reverse(inbytes)
	fmt.Println(string(inbytes))
}

func reverse(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}
