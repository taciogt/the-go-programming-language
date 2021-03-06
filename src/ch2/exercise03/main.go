package main

import "fmt"

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

func main() {
	fmt.Println(pc)
	for i := uint64(0); i <= 32; i++ {
		fmt.Printf("Popcount(%d): %d\n", i, PopCount(i))
		fmt.Printf("LoopedPopcount(%d): %d\n", i, LoopedPopCount(i))
	}
}
