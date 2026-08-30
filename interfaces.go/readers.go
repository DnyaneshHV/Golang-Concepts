package main

import (
    "fmt"
    "strings"
)

func main() {
    reader := strings.NewReader("Hello")  // suppose the book contains "Hellow"

    buf := make([]byte, 5)   // create the buffer with the 5 spaces 

    n, err := reader.Read(buf) // reader copies the text into the box 

    fmt.Println(n)
    fmt.Println(string(buf))
    fmt.Println(err)
}