package main

import "fmt"

type Vertex struct {
        Lat, Long float64
}

var m = map[string]Vertex{
        "hoge": Vertex {
                1, 2,
        },
        "foo" : Vertex {
                3, 4,
        },
}

func main() {
        fmt.Println(m)
}
