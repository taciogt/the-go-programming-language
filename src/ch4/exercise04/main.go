package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	ch1 := sha256.Sum256([]byte("x"))
	ch2 := sha256.Sum256([]byte("X"))

	diffArray := [32]byte{}
	popCount := [32]int{}
	var totalDiffs int
	for i := 0; i < 32; i++ {
		diffArray[i] = ch1[i] ^ ch2[i]
		popCount[i] = PopCount(uint(diffArray[i]))
		totalDiffs += popCount[i]
	}

	fmt.Printf("ch1: %08b\n"+
		"ch2: %08b\n"+
		"diff:%08b\n"+
		"popCount: %d\n"+
		"totalDiffs: %d\n",
		ch1, ch2, diffArray, popCount, totalDiffs)
}

func PopCount(x uint) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
