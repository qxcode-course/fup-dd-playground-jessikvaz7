
package main

import "fmt"

func main() {
	var velocidade, tempo, consumo float64
	var desempenho float64

	fmt.Scan(&velocidade)
	fmt.Scan(&tempo)
	fmt.Scan(&consumo)

	tempo = tempo / 60
	desempenho = (velocidade * tempo) / consumo

	fmt.Printf("%.2f\n", desempenho)
}