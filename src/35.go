package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Print("Enter your name: ")
    var name string
    _, err := fmt.Scanln(&name)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    fmt.Printf("Hello, %s!\n", name)
}
