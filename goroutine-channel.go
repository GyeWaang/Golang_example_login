package main

import (
    "fmt"
)

func producer(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}

func consumer(ch chan int) {
    for i := range ch {
        fmt.Println("Received:", i)
    }
}

func main() {
    ch := make(chan int)
    go producer(ch)
    go consumer(ch)

    var input string
    fmt.Scanln(&input)
}
