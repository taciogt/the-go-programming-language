// How to run the benchmarks: go test -bench=. ./exercise03
package main

import (
	"fmt"
	"math"
	"testing"
)

var runCount uint64

func init() {
	runCount = uint64(math.Pow(2, 16))
	fmt.Println(runCount)
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint64(0); j < runCount; j++ {
			PopCount(j)
		}
	}
}

func BenchmarkLoopedPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint64(0); j < runCount; j++ {
			LoopedPopCount(j)
		}
	}
}

func BenchmarkPopCount64Shift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint64(0); j < runCount; j++ {
			PopCount64Shift(j)
		}
	}
}

func BenchmarkTakeRightMost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint64(0); j < runCount; j++ {
			TakeRightMost(j)
		}
	}
}
