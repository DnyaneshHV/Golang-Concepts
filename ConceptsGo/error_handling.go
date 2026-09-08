package main

import (
    "errors"
    "fmt"
)

// error_handling.go
// Demonstrates basic error creation and handling.

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    if res, err := divide(10, 2); err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10/2 =", res)
    }

    if _, err := divide(10, 0); err != nil {
        fmt.Println("Expected error:", err)
    }
}
