package main

import "fmt"

func main() {
	var nota1, nota2, notaFinal int
	var media float64

	fmt.Scan(&nota1)
	fmt.Scan(&nota2)
	fmt.Scan(&notaFinal)

	media = float64(nota1+nota2) / 2

	if media >= 7 {
		fmt.Println("aprovado")
	} else if media < 4 {
		fmt.Println("reprovado")
	} else {
		media = (media + float64(notaFinal)) / 2

		if media >= 5 {
			fmt.Println("aprovado na final")
		} else {
			fmt.Println("reprovado na final")
		}
	}
}