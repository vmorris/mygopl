// Copyright © 2017 Alan A. A. Donovan, Brian W. Kernighan & Vance E. Morris
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func Example_one() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // "3"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())    // "2"

	y.UnionWith(&x)
	fmt.Println(y.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	y.Remove(42)
	fmt.Println(y.String()) // "{1 9 144}"

	y.Clear()
	fmt.Println(y.String()) // "{}"

	z := x.Copy()
	fmt.Println(z.String()) // "{1 9 144}"

	z.AddAll(2, 3)
	fmt.Println(z.String()) // "{1 2 3 9 144}"

	y.Clear()
	x.Clear()
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5, 6)
	x.IntersectWith(&y)
	fmt.Println(x.String()) // "{3 4}"

	y.Clear()
	x.Clear()
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 5, 6, 128)
	x.IntersectWith(&y)
	fmt.Println(x.String()) // "{3}"

	y.Clear()
	x.Clear()
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5, 6)
	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{1 2}"
	y.Clear()
	x.Clear()
	x.AddAll(1, 2, 3, 4)
	y.AddAll(4, 5, 2)
	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{1 3}"
	y.Clear()
	x.Clear()
	x.AddAll(4, 3)
	y.AddAll(4, 5, 2)
	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{3}"

	y.Clear()
	x.Clear()
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5, 6)
	x.SymmetricDifference(&y)
	fmt.Println(x.String()) // "{1 2 5 6}"

	// Output:
	// {1 9 144}
	// 3
	// {9 42}
	// 2
	// {1 9 42 144}
	// true false
	// {1 9 144}
	// {}
	// {1 9 144}
	// {1 2 3 9 144}
	// {3 4}
	// {3}
	// {1 2}
	// {1 3}
	// {3}
	// {1 2 5 6}

}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
