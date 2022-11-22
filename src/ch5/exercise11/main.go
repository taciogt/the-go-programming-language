package main

import (
	"fmt"
	"log"
)

var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"linear algebra":        {"calculus"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro do programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	result, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i, course := range result {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(key string, branch map[string]bool) error
	visitAll = func(k string, branch map[string]bool) error {
		if _, ok := branch[k]; !ok {
			branch[k] = true
		} else {
			return fmt.Errorf("cycle 1")
		}

		for _, item := range m[k] {
			if err := visitAll(item, branch); err != nil {
				return err
			}
			if !seen[item] {
				seen[item] = true
				order = append(order, item)
			}

		}
		if !seen[k] {
			seen[k] = true
			order = append(order, k)
		}
		delete(branch, k)
		return nil
	}
	for k := range m {
		if err := visitAll(k, make(map[string]bool)); err != nil {
			return order, err
		}
	}

	return order, nil
}
