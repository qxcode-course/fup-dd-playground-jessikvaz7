package main

import "fmt"

func main() {
    var s string
     fmt.Scanln(&s)
     c := rune(s[0])

    if c >= 'a' && c <= 'z' {
     fmt.Printf("%c\n", c - 32)
    } else if c >= 'A' && c <= 'Z' {
        fmt.Printf("%c\n", c + 32)
    } else {
        fmt.Printf("%c\n", c)
    }
}