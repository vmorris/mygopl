// printsha prints the SHA256 hash of stdin by default and allows cli flag
// switches to print SHA384 or SHA512 hash instead

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	sha384Ptr := flag.Bool("sha384", false, "print SHA384 hash")
	sha512Ptr := flag.Bool("sha512", false, "print SHA512 hash")

	flag.Parse()
	//fmt.Printf("printsha: sha384Ptr: %t, sha512Ptr: %t\n", *sha384Ptr, *sha512Ptr)

	if *sha384Ptr && *sha512Ptr {
		fmt.Fprintln(os.Stderr, "You cannot use both sha384 and sha512 options together!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Print(err)
	}

	switch {
	case *sha384Ptr:
		doSHA384(bytes)
	case *sha512Ptr:
		doSHA512(bytes)
	default:
		doSHA256(bytes)
	}

}

func doSHA256(b []byte) {
	h := sha256.New()
	h.Write([]byte(b))
	fmt.Printf("%x\n", h.Sum(nil))
}

func doSHA384(b []byte) {
	h := sha512.New384()
	h.Write([]byte(b))
	fmt.Printf("%x\n", h.Sum(nil))
}

func doSHA512(b []byte) {
	h := sha512.New()
	h.Write([]byte(b))
	fmt.Printf("%x\n", h.Sum(nil))
}
