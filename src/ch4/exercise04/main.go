package main

import "fmt"

func main() {
	for step := 0; step < 5; step++ {
		a := []int{0, 1, 2, 3, 4, 5}
		rotate(a, step)
		fmt.Printf("steps: %d, a: %d\n", step, a)
	}
}

func rotate(s []int, steps int) {
	begin := make([]int, len(s)-steps)
	copy(begin, s[steps:])
	end := make([]int, steps)
	copy(end, s[:steps])

	for i, v := range begin {
		s[i] = v
	}
	for i, v := range end {
		nextIndex := len(s) - steps + i
		s[nextIndex] = v
	}

}
