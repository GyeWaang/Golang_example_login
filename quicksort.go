// gyewong
package main

import "fmt"

func quicksort(a []int) []int {
    if len(a) <= 1 {
        return a
    }

    pivot := a[len(a)/2]
    var less []int
    var greater []int

    for _, x := range a {
        if x <= pivot {
            less = append(less, x)
        } else {
            greater = append(greater, x)
        }
    }

    return append(quicksort(less), pivot, quicksort(greater)...)
}

func main() {
    a := []int{4, 2, 1, 7, 3, 6, 5}
    fmt.Println(quicksort(a))
}
