package main

import "fmt"

func main() {
        p := []int{2, 3, 5, 7, 11, 13}
        q := p[1:3]
        fmt.Println("p == ", p)
        fmt.Println("q == ", q)
        q[1] = 2
        fmt.Println("p == ", p)
        fmt.Println("q == ", q)
}
