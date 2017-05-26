package popcount

/*
vance@holocron:~/go/src/github.com/vmorris/mygopl/ch2/ex2.3/popcount$ go test -bench=.
BenchmarkPopCount-4             	2000000000	         0.25 ns/op
BenchmarkPopCount2-4            	100000000	        17.4 ns/op
BenchmarkBitCount-4             	2000000000	         0.25 ns/op
BenchmarkPopCountByClearing-4   	100000000	        16.0 ns/op
BenchmarkPopCountByShifting-4   	30000000	        54.7 ns/op
PASS
ok  	github.com/vmorris/mygopl/ch2/ex2.3/popcount	6.144s
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

// PopCount2 returns the population count (number of set bits) of x, using a loop
func PopCount2(x uint64) int {
	var sum byte
	for i := uint(0); i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}
