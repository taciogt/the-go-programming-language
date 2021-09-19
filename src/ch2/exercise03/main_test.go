// How to run the benchmarks: go test -bench=. ./exercise03
package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint64(0); j < 256; j++ {
			PopCount(j)
		}
	}
}

func BenchmarkLoopedPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint64(0); j < 256; j++ {
			LoopedPopCount(j)
		}
	}
}
