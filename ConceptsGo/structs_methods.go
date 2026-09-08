package main

import "fmt"

// structs_methods.go
// Demonstrates structs and methods, and the difference between value and pointer receivers.

type Counter struct{ n int }

// Value receiver: does not modify original when called on a value.
func (c Counter) Value() int { return c.n }

// Pointer receiver: modifies the original when called on a pointer.
func (c *Counter) Increment() { c.n++ }

func main() {
    c := Counter{}
    c.Increment()
    fmt.Println("Value after increment:", c.Value())
    p := &c
    p.Increment()
    fmt.Println("Value after pointer increment:", c.Value())
}
