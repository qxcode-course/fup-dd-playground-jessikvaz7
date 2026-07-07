package main

import "fmt"

type Jogada struct {
	p1, p2 int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// retorna se a jogada é válida e a pontuação
func calcPontuacao(j Jogada) (bool, int) {
	if j.p1 < 10 || j.p2 < 10 {
		return false, 0
	}
	return true, abs(j.p1-j.p2)
}

func procurarMelhorJogada(jogadas []Jogada) int {
	melhor := -1
	menorPont := 0

	for i := 0; i < len(jogadas); i++ {
		valida, pont := calcPontuacao(jogadas[i])

		if !valida {
			continue
		}

		if melhor == -1 || pont < menorPont {
			melhor = i
			menorPont = pont
		}
	}

	return melhor
}

func main() {
	var n int
	fmt.Scan(&n)

	jogadas := make([]Jogada, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&jogadas[i].p1, &jogadas[i].p2)
	}

	resp := procurarMelhorJogada(jogadas)

	if resp == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(resp)
	}
}