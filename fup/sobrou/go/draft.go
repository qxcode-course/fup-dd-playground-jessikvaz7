package main

import "fmt"

func main() {
    var q1, q2, q3 int
    var v1, v2, v3 float64
    var dinheiro float64

    // Entrada
    fmt.Scan(&q1)
    
fmt.Scan(&q2)
    fmt.Scan(&q3)
    fmt.Scan(&v1)
    fmt.Scan(&v2)
    fmt.Scan(&v3)

    fmt.Scan(&dinheiro)

    // Cálculo
    total := float64(q1)*v1 + float64(q2)*v2 + float64(q3)*v3
    troco := dinheiro - total

    // Saída formatada
    fmt.Printf("%.2f\n", troco)
}