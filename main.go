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

type User struct {
    Name string
    Age  int
}

func main() {
    user := User{
        Name: "Bob",
        Age:  18,
    }
    showName(&user)

}

func showName(user *User) {
    fmt.Println(user.Name)
}
