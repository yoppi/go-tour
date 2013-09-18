package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v2 := Vertex{3, 4}
    v := &Vertex{3, 4}
    fmt.Println(v)
    fmt.Println(v2)
    fmt.Println(v.Abs())
    fmt.Println(v2.Abs())
}
