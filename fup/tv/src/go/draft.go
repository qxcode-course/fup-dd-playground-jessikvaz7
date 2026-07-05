package main
import "fmt"
func main() {

    var valor float64
    var parcelas int
    
    fmt.Scan(&valor)
    fmt.Scan(&parcelas)

    juros := (parcelas - 1) * 5

    total := valor + valor*float64(juros)/100
    valorParcela := total / float64(parcelas)

    fmt.Printf("%.2f\n", valorParcela)
    fmt.Printf("%.2f\n", total)
}