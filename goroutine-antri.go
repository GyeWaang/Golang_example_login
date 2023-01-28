package main

import (
    "fmt"
)

func main() {
    // Create a channel to signal completion
    done := make(chan bool)

    // Send the first request
    go func() {
        fmt.Println("asu1")
        fmt.Println("asu2")
        fmt.Println("asu3")
        fmt.Println("asu4")
        fmt.Println("asu5")
        fmt.Println("asu6")
        done <- true
    }()

    // Send the second request
    go func() {
        fmt.Println("asua")
        fmt.Println("asub")
        fmt.Println("asuc")
        fmt.Println("asud")
        fmt.Println("asue")
        fmt.Println("asuf")
        done <- true
    }()

    // Wait for both functions to complete
    <-done
    <-done
}
