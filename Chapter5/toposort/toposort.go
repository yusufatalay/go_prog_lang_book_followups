package main

import (
	"fmt"
)

// prereqs map computer science courses to their prerequisites
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	var stackorder int
	// the map is now string->int instead of string->bool to to expand the
	// function of the map to see if there is a cyclicity in it.
	seen := make(map[string][]int)
	// gotta declare the anon function first then assign it for recurrency
	// purposes
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			// if item havent been visited before
			if seen[item][0] == 0 {
				seen[item][0] = 1
				seen[item][1] = stackorder
				stackorder++
				visitAll(m[item])
				order = append(order, item)
				// if the item has visited multiple time
			}
			if seen[item][1] > 0 {
				fmt.Printf("Cyclicity detected at %s", item)
			}

		}
	}

	//	var keys []string
	//	for key := range m {
	//		keys = append(keys, key)
	//	}
	//sort.Strings(keys)
	for key := range m {
		visitAll([]string{key})
	}
	return order

}
