package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
        ret := make([][]uint8, dy)
        for i:=0; i<len(ret); i++ {
                ret[i] = make([]uint8, dx)
        }
        return ret
}

func main() {
        pic.Show(Pic)
}
