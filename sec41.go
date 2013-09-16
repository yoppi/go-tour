package main

import (
        "code.google.com/p/go-tour/wc"
        "strings"
)

func WordCount(s string) map[string]int {
        ret := make(map[string]int)
        for _, word := range strings.Fields(s) {
                if _, ok := ret[word]; ok {
                        ret[word] += 1
                } else {
                        ret[word] = 1
                }
        }
        return ret
}

func main() {
        wc.Test(WordCount)
}
