package main

import "testing"


func BenchmarkConcatenateWithSimpleFor(b *testing.B){
	for i := 0; i < b.N; i++ {
		ConcatenateWithSimpleFor([]string{"filename", "Some", "string", "benchmark"})
	}
}

func BenchmarkConcatenateWithForInRange(b *testing.B){
	for i := 0; i < b.N; i++ {
		ConcatenateWithForInRange([]string{"filename", "Some", "string", "benchmark"})
	}
}

func BenchmarkConcatenateWithJoin(b *testing.B){
	for i := 0; i < b.N; i++ {
		ConcatenateWithJoin([]string{"filename", "Some", "string", "benchmark"})
	}
}