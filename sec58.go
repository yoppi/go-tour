package main

import  (
  "net/http"
  "fmt"
)

type String string
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

type Struct struct {
    A string
    B string
    C string
}
func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func main() {
    http.Handle("/string", String("Hoge"))
    http.Handle("/struct", &Struct{"Hello", ":", "Gopher"})
    http.ListenAndServe("localhost:3001", nil)
}
