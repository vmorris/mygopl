package main

import "fmt"

func max(vals ...int) (result int, ok bool) {
	if len(vals) == 0 {
		return
	}
	ok = true
	result = vals[0]
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return
}

func min(vals ...int) (result int, ok bool) {
	if len(vals) == 0 {
		return
	}
	ok = true
	result = vals[0]
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return
}

func main() {

	values := []int{1, 2, 3, 4, 5, 6, 7, 8}

	if result, ok := max(values...); ok {
		fmt.Println("max: ", result)
	} else {
		fmt.Println("not ok")
	}

	if result, ok := min(values...); ok {
		fmt.Println("min: ", result)
	} else {
		fmt.Println("not ok")
	}

	// can't call max() or min() without at least 1 arg..

	if result, ok := max(); ok {
		fmt.Println("max: ", result)
	} else {
		fmt.Println("not ok")
	}

	if result, ok := min(); ok {
		fmt.Println("min: ", result)
	} else {
		fmt.Println("not ok")
	}
}
