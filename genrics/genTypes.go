package main

import "fmt"

// List represents a singly-linked list node
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	// Create nodes
	node3 := &List[int]{val: 30}
	node2 := &List[int]{val: 20, next: node3}
	node1 := &List[int]{val: 10, next: node2}

	// Traverse and print
	for node := node1; node != nil; node = node.next {
		fmt.Println(node.val)
	}
}