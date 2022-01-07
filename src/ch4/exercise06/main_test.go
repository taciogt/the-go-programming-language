package main

import (
	"fmt"
	"testing"
)

func Test_squashSpaces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc", "abc"},
		{"aabac", "aabac"},
		{"aa bac", "aa bac"},
		{"aa ba c", "aa ba c"},
		{"aa ba   c", "aa ba c"},
		{"aa ba  \n c", "aa ba c"},
		{"aa\t\n ba  \n c", "aa ba c"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("testing for %s", tt.input), func(t *testing.T) {
			originalInput := make([]byte, len(tt.input))
			copy(originalInput, tt.input)

			input := []byte(tt.input)
			squashSpaces(&input)
			if string(input) != tt.expected {
				t.Fatalf(`squashSpaces("%v") -> %v, want "%s"`, originalInput, input, tt.expected)
			}
		})
	}
}
