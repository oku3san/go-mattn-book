package main

import "fmt"

func main() {

    message := "hi"
    go sendMessage(message)
}

func sendMessage(msg string) {
    fmt.Println(msg)
}
