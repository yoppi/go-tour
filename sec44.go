package main

import "fmt"

func fib() func() int {
    x0 := 0
    x1 := 1
    xn := x0 + x1
    return func() int {
        x0 = x1
        x1 = xn
        xn = x0 + x1
        return xn
    }
}

func main() {
    f := fib()
    for i:= 0; i<10; i++ {
        fmt.Println(f())
    }
}
