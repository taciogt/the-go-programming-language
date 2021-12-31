package main

import "fmt"

type stack []int

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("s: %d\n", s)

	s = push(s, 4)
	fmt.Printf("s: %d\n", s)

	s = remove(s, 2)
	fmt.Printf("s: %d\n", s)
}

func push(s stack, v int) stack {
	return append(s, v)
}

func remove(s stack, i int) stack {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
