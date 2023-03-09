package main

import "fmt"

//type Fruit int
//type Animal int
//
//const (
//    Apple Fruit = iota
//    Orange
//    Banana
//)
//
//const (
//    Monkey Animal = iota
//    Elephant
//    Pig
//)

func main() {
    //var fruit Fruit = Apple
    //fmt.Println(fruit)
    //
    //fruit = Elephant
    //fmt.Println(fruit)

    //a := []int{1, 2, 3}
    //fmt.Println(a)
    //a = append(a, 4, 5, 6)
    //fmt.Println(a, len(a))
    //
    //fmt.Println(a[2:5])

    //s := "こんにちわ世界"
    //rs := []rune(s)
    //rs[4] = 'は'
    //s = string(rs)
    //println(s)

    m := make(map[string]int)
    m["John"] = 21
    m["Bob"] = 18
    m["Mark"] = 33

    fmt.Println(m)
}
