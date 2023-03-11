package main

import "fmt"

type Value int

func (v *Value) Add(n Value) {
    //return v + n
    *v += n
}

func main() {
    v := Value(1)
    v.Add(2)
    fmt.Println(v)

    //v := 1
    //p := &v
    //*p = 2
    //fmt.Println(v)
}
