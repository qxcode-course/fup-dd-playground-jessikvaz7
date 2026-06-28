package main
import "fmt"
func main() {
    var a, b int

    fmt.Scan(&a)
    fmt.Scan(&b)

    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Printf("%.2f\n", float64(a) / float64(b) )
    fmt.Println(a % b)
}