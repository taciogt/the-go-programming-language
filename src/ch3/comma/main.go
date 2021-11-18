package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 12, 123, 1234, 12345, 123456, 1234567}
	for _, n := range nums {
		fmt.Printf("%d: %s\n", n, comma(fmt.Sprintf("%d", n)))
	}

	fStrings := []string{"1", "12", "123.45", "1222.3333"}
	for _, n := range fStrings {
		fmt.Printf("%s: %s\n", n, commaForFloat(n))
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

func commaForFloat(s string) string {
	i := strings.Index(s, ".")
	if i == -1 {
		return comma(s)
	}
	return comma(s[:i]) + s[i:]
}
