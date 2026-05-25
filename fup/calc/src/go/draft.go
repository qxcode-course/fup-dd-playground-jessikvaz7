package main

import "fmt"

func main() {

    var n1 int
    var n2 int
    var S string

    fmt.Scanln(&n1)
    fmt.Scanln(&n2)
    fmt.Scanln(&S)

    switch S {
    case "+":
        fmt.Println(n1 + n2)
    case "-":
        fmt.Println(n1 - n2)
    case "*":
        fmt.Println(n1 * n2)
    case "/":
        fmt.Println(n1 / n2)
    }
}