package main

import "fmt"

func main() {
    var a, b, c int
    
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    
    if (a > b && a < c) || (a > c && a < b) {
        fmt.Println(a)
    } else if (b > a && b < c) || (b > c && b < a) {
        fmt.Println(b)
    } else {
        fmt.Println(c)
    }     
}