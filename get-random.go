package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    squadlist := []string{"item1", "item2", "item3", "item4", "item5"}
    randomIndex := rand.Intn(len(squadlist))
    randomItem := squadlist[randomIndex]
    fmt.Println(randomItem)
}
