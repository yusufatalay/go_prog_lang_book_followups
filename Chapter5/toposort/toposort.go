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
	//	"linear algebra":        {"calculus"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	indegree := make(map[string]int)
	var queue []string
	// gotta declare the anon function first then assign it for recurrency
	// visitAll is an anonymous-recursive function that visits each value of
	// given key
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			// if item havent been visited before
			if !seen[item] {
				seen[item] = true
				indegree[item] = len(m[item])
				if indegree[item] == 0 {
					queue = enqueue(queue, item)
				}
				visitAll(m[item])
				order = append(order, item)
			}

		}
	}

	for key := range m {

		visitAll([]string{key})
	}
	var viscount int
	var poppedelement string
	for len(queue) != 0 {
		queue, poppedelement = dequeue(queue)
		for _, el := range m[poppedelement] {
			indegree[el]--

			if indegree[el] == 0 {
				queue = enqueue(queue, el)
			}
		}
		viscount++
	}
	if viscount != len(order) {
		fmt.Println(viscount, len(order))
		fmt.Println("there is a cycle")
	}
	return order
}

func contains(a []string, s string) bool {
	for _, e := range a {
		if e == s {
			return true
		}
	}
	return false
}

func enqueue(queue []string, element string) []string {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue []string) ([]string, string) {
	return queue[1:], queue[0] // Slice off the element once it is dequeued.
}
