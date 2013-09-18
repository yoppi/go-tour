package main

import (
    "fmt"
    "math"
)

type Equaler interface {
    Equal(Equaler) bool
}

// uはEqualerではなくてあくまでT型なのでEqualerを満たしていない
type T int
func (t T) Equal(u T) bool { return t == u }

type T2 int
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }

type Abser interface {
    Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func main() {
    var a Abser
    fmt.Println(a) //=> <nil>

    f := MyFloat(-math.Sqrt2)
    fmt.Println(f)

    v := Vertex{3, 4}
    fmt.Println(v)

    a = f // 代入することでimplementsをコンパイル時に判定できる
    fmt.Println(a.Abs())

    t := T(1)
    fmt.Println(t.Equal(1)) // Tじゃない普通の1を渡せる

    t2 := T2(1)
    //fmt.Println(t2.Equal(1)) // Equalerを見たさないものは渡せない
    fmt.Println(t2.Equal(T2(1))) // ただし、Equalerを満たすものなら何でも渡せるので結局のところ実行時に型チェックが必要になる
}


