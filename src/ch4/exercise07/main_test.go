package main

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		word []byte
	}
	tests := []struct {
		input    string
		expected string
	}{
		{"abcdef", "fedcba"},
		{"abc", "cba"},
		{"aba", "aba"},
		{"abab", "baba"},
		{"abxz", "zxba"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("testing for %s", tt.input), func(t *testing.T) {
			originalInput := string(tt.input)
			bs := []byte(tt.input)
			reverse(bs)
			if string(bs) != tt.expected {
				t.Fatalf(`reverse("%s") -> %s, want "%s"`, originalInput, bs, tt.expected)
			}

		})
	}
}
