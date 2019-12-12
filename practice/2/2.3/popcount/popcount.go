package popcount

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func TimeConsuming(tag string) func() {
	now := time.Now().UnixNano()
	return func() {
		after := time.Now().UnixNano()
		fmt.Printf("%q time cost %d ns\n", tag, after-now)
	}
}

func PopCount1(x uint64) int {
	defer TimeConsuming("PopCount1")()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	defer TimeConsuming("PopCount2")()
	var n int
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}
