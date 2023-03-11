package main

import "fmt"

type Value int

func (v Value) Add(n Value) Value {
    return v + n
}

func main() {
    v := Value(1)
    v = v.Add(2)
    fmt.Println(v)
}
