package main

import "fmt"

func main() {
        a := make([]int, 5)
        fmt.Println(a, cap(a), len(a))

        b := make([]int, 5, 1000)
        fmt.Println(b, len(b), cap(b))
}
