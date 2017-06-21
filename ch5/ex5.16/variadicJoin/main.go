package main

import "fmt"

// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
// To make this variadic, sep must come first in the arg list.
func Join(sep string, a ...string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1]
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1] + sep + a[2]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func main() {

	strs := []string{"This", "is", "a", "bunch", "of", "strings"}

	fmt.Println(Join("-", strs...))
	fmt.Println(Join("-"))
	fmt.Println(Join("-", "This", "is", "some", "more", "strings"))

}

// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func originalJoin(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1]
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1] + sep + a[2]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
