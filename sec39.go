package main

import "fmt"

type Vertex struct {
        Lat, Long float64
}

var m = map[string]Vertex{
        "Bell Labs": {1, 2},
        "Google": {3, 4},
}

func main() {
        fmt.Println(m)
}
