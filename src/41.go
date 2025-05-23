package main

import (
    "fmt"
)

func main() {
    // Example of a simple Go program to demonstrate basic operations.
    var a int = 5;
    var b int = 3;

    // Addition
    sum := a + b;
    fmt.Println("The sum is:", sum);

    // Subtraction
    sub := a - b;
    fmt.Println("The subtraction is:", sub);

    // Multiplication
    mul := a * b;
    fmt.Println("The multiplication is:", mul);

    // Division (integer division)
    div := a / b;
    if div == int(div) {
        fmt.Println("Division result is an integer.")
    } else {
        fmt.Println("Division result is not an integer.")
    }
}
