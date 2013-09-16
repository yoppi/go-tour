package main

import "fmt"

// 関数外では変数宣言に必ずvarが必要
var x, y, z = 1, 2, 3
// non-declaraiton statement outside function body
//foo, bar, baz := 1, 2, 3

func main() {
        c, python, java := true, false, "!"
        fmt.Println(x, y, z)
        fmt.Println(c, python, java)
}

