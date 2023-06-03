package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
    fahrenheit := (celsius * 9 / 5) + 32
    return fahrenheit
}

func main() {
    celsius := 25.0
    fahrenheit := celsiusToFahrenheit(celsius)
    fmt.Printf("%.2f degrees Celsius is equal to %.2f degrees Fahrenheit\n", celsius, fahrenheit)
}

