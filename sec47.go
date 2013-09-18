package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("morning")
    case t.Hour() < 17:
        fmt.Println("afternoon")
    default:
        fmt.Println("evening")
    }
}
