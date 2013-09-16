package main

import "fmt"

func main() {
        m := make(map[string]int)

        m["answer"] = 42
        fmt.Println(m["answer"])

        delete(m, "answer")
        fmt.Println(m)

        m["answer"] = 1
        v, ok := m["answer"]
        fmt.Println(v, ok)
}
