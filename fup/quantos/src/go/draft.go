package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    fmt.Scanln(&c)

    if a == b && b == c {
        fmt.Println(3)
} else if a == b || a == c || b == c {
       fmt.Println(2)
 } else {
       fmt.Println(0)

  }
}