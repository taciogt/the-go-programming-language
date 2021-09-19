package main

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	fmt.Println(pc)
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
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

// LoopedPopCount returns the population count (number of set bits) of x
func LoopedPopCount(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		b += pc[byte(x>>(i*8))]
	}
	return int(b)
}

func PopCount64Shift(x uint64) int {
	var c int
	for i := 0; i < 64; i++ {
		c += int(x >> i & 1)
	}
	return c
}

func main() {
	fmt.Println(pc)
	for i := uint64(0); i <= 255; i++ {
		fmt.Printf("%3d = %08b | Popcount: %d | LoopedPopcount: %d | PopCount64Shift: %d\n", i, i, PopCount(i), LoopedPopCount(i), PopCount64Shift(i))
	}
}
