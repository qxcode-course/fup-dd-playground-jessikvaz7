package main

import "fmt"

type Restaurante struct {
	nome  string
	ponto int
}

func melhorRestaurante(lista []Restaurante) string {
	melhor := lista[0]

	for i := 1; i < len(lista); i++ {
		if lista[i].ponto > melhor.ponto ||
			(lista[i].ponto == melhor.ponto && lista[i].nome < melhor.nome) {
			melhor = lista[i]
		}
	}

	return melhor.nome
}

func main() {
	var n int
	fmt.Scan(&n)

	lista := make([]Restaurante, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&lista[i].nome, &lista[i].ponto)
	}

	fmt.Println(melhorRestaurante(lista))
}