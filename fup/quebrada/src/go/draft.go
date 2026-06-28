package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    fmt.Println(n1 / n2)
    fmt.Println(n1 % n2)
    fmt.Printf("%.2f\n", float64(n1)/float64(n2))

}