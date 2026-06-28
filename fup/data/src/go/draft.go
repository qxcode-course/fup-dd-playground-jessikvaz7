package main
import "fmt"
func main() {
    var hora, minuto, dia, mes, ano int

    fmt.Scan(&hora)
    fmt.Scan(&minuto)
    fmt.Scan(&dia)
    fmt.Scan(&mes)
    fmt.Scan(&ano)

    fmt.Printf("%02d:%02d %02d/%02d/%02d\n", hora, minuto, dia, mes, ano%100)
}