package main

import "fmt"

type Vertex struct {
        X int
        Y int
}

func main() {
        p := Vertex{1, 2}
        q := &p
        q.X = 1e9
        p.Y = 10
        fmt.Println(p)
}
