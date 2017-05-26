// convert numeric arguments between various units.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vmorris/mygopl/ch2/ex2.1/tempconv"
	"github.com/vmorris/mygopl/ch2/ex2.2/lengthconv"
)

func main() {

	var l []string

	if len(os.Args) == 1 {
		s := bufio.NewScanner(os.Stdin)
		fmt.Println("Reading a single line from Stdin...")
		s.Scan()
		in := s.Text()
		fmt.Sprintln(in)
		if len(in) == 0 {
			os.Exit(0) // empty Stdin
		}
		l = strings.Split(in, " ")
	} else {
		l = os.Args[1:]
	}

	doConversions(l)
}

func doConversions(l []string) {

	for _, arg := range l {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		ft := lengthconv.Feet(t)
		m := lengthconv.Meters(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), f, tempconv.FToK(f))
		fmt.Printf("%s = %s, %s = %s\n",
			c, tempconv.CToF(c), c, tempconv.CToK(c))
		fmt.Printf("%s = %s, %s = %s\n",
			k, tempconv.KToF(k), k, tempconv.KToC(k))
		fmt.Printf("%s = %s, %s = %s\n",
			ft, lengthconv.FToM(ft), m, lengthconv.MToF(m))
	}
}
