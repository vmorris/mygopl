package popcount

/*
vance@holocron:~/go/src/github.com/vmorris/mygopl/ch2/ex2.5/popcount$ go test -bench=.
BenchmarkPopCount-4    	2000000000	         0.23 ns/op
BenchmarkPopCount4-4   	100000000	        13.2 ns/op
PASS
ok  	github.com/vmorris/mygopl/ch2/ex2.5/popcount	1.829s
*/

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount4 clears the rightmost non-zero bit of x while counting population
func PopCount4(x uint64) int {
	var sum byte
	for x != 0 { // if not zero, clear and add 1
		x = x & (x - 1)
		sum++
	}
	return int(sum)
}
