package main
import (
    "fmt"
    "math"
)
func main() {
    var op string
    var num float64

    fmt.Scan(&op)
    fmt.Scan(&num)

    switch op {
    case "c":
        fmt.Println(int(math.Ceil(num))) 
    case "f":
        fmt.Println(int(math.Floor(num)))
    case "r":
        fmt.Println(int(math.Round(num)))
    }            
}