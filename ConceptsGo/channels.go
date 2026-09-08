package main

import (
    "fmt"
    "time"
)

// channels.go
// Demonstrates channels with a simple producer/consumer pattern.

func producer(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i
        fmt.Println("Produced", i)
    }
    close(ch)
}

func consumer(ch <-chan int, done chan<- struct{}) {
    for v := range ch {
        fmt.Println("Consumed", v)
        time.Sleep(50 * time.Millisecond)
    }
    done <- struct{}{}
}

func main() {
    ch := make(chan int)
    done := make(chan struct{})
    go producer(ch)
    go consumer(ch, done)
    <-done
    fmt.Println("Processing complete")
}
