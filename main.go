package main

import "fmt"

func main() {
    //defer fmt.Println("6")
    //defer fmt.Println("5")
    //defer fmt.Println("4")
    //fmt.Println("1")
    //fmt.Println("2")
    //fmt.Println("3")
    doSomething()
}

func doSomething() {
    var n = 1
    defer func() {
        fmt.Println(n)
    }()

    n = 2
}
