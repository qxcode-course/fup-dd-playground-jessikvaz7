package main
import "fmt"
func main() {
    var c float64
    fmt.Scan(&c)

    f := (c * 1.8) + 32
    fmt.Printf("%.6f\n", f)

}
