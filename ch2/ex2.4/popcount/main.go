package popcount

/*
vance@holocron:~/go/src/github.com/vmorris/mygopl/ch2/ex2.4/popcount$ go test -bench=.
BenchmarkPopCount-4    	2000000000	         0.23 ns/op
BenchmarkPopCount3-4   	50000000	        34.1 ns/op
PASS
ok  	github.com/vmorris/mygopl/ch2/ex2.4/popcount	2.232s
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

// PopCount3 returns the population count (number of set bits) of x, using a loop
// that bitshifts through 64 positions, testing the rightmost bit
func PopCount3(x uint64) int {
	var sum byte
	for i := uint(0); i < 64; i++ {
		sum += byte(x & 1)
		x = x >> 1
	}
	return int(sum)
}
