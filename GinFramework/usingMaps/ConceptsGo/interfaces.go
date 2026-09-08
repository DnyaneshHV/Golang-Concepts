package main

import "fmt"

// interfaces.go
// Demonstrates Go interfaces and polymorphism.

type Speaker interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow" }

func saySomething(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    var s Speaker
    s = Dog{}
    saySomething(s)
    s = Cat{}
    saySomething(s)
}
