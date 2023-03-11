package main

import "fmt"

func main() {
    defer fmt.Println("6")
    defer fmt.Println("5")
    defer fmt.Println("4")
    fmt.Println("1")
    fmt.Println("2")
    fmt.Println("3")
}
