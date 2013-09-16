package main

import "fmt"

func main() {
        m := make(map[string]int)

        m["answer"] = 42
        fmt.Println(m["answer"])

        delete(m, "answer")
        fmt.Println(m)

        m["answer"] = 1
        if v, ok := m["answer"]; ok {
                // v使う
                fmt.Println(v, ok)
        } else {
                // ない場合の処理
        }
}
