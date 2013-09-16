package main

import (
        "fmt"
        "math/rand"
        "time"
)

func main() {
        time := time.Now().UnixNano()
        rand.Seed(time)
        fmt.Println("My favorite number is", rand.Intn(100))
}
