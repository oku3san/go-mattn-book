package main

import "fmt"

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
