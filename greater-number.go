package main

import "fmt"

func main() {
    var num1, num2 int

    fmt.Print("Enter the first number: ")
    fmt.Scanln(&num1)

    fmt.Print("Enter the second number: ")
    fmt.Scanln(&num2)

    if num1 > num2 {
        fmt.Println("The first number is greater.")
    } else if num2 > num1 {
        fmt.Println("The second number is greater.")
    } else {
        fmt.Println("Both numbers are equal.")
    }
}

