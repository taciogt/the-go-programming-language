package main

import "fmt"

func main() {
	nums := []int{1, 12, 123, 1234, 12345, 123456, 1234567}
	for _, n := range nums {
		fmt.Printf("%d: %s\n", n, comma(fmt.Sprintf("%d", n)))
	}
}

// comma inserts commas in a non-negative decimal integer string
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
