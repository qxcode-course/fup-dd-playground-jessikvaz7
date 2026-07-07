package main

import (
	"fmt"
	"math/rand"
	"time"
)

func nome(op int) string {
	switch op {
	case 0:
		return "PEDRA"
	case 1:
		return "PAPEL"
	case 2:
		return "TESOURA"
	case 3:
		return "LAGARTO"
	default:
		return "SPOCK"
	}
}

func ganhou(a, b int) bool {
	return (a == 0 && (b == 2 || b == 3)) ||
		(a == 1 && (b == 0 || b == 4)) ||
		(a == 2 && (b == 1 || b == 3)) ||
		(a == 3 && (b == 4 || b == 1)) ||
		(a == 4 && (b == 2 || b == 0))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		voce, pc := 0, 0

		for round := 1; round <= 5; round++ {
			fmt.Println("# JOKENPÔ #")
			fmt.Printf("Você: %d | PC: %d\n", voce, pc)
			fmt.Printf("Round: %d / 5\n\n", round)

			fmt.Println("1 - Pedra")
			fmt.Println("2 - Papel")
			fmt.Println("3 - Tesoura")
			fmt.Println("4 - Lagarto")
			fmt.Println("5 - Spock")

			var op int
			fmt.Scan(&op)

			if op < 1 || op > 5 {
				fmt.Println("Jogada inválida!")
				round--
				continue
			}

			jogador := op - 1
			computador := rand.Intn(5)

			fmt.Printf("Você jogou %s e o PC %s.\n", nome(jogador), nome(computador))

			if jogador == computador {
				fmt.Println("Ninguém ganhou!")
			} else if ganhou(jogador, computador) {
				fmt.Println("Você ganhou!")
				voce++
			} else {
				fmt.Println("O PC ganhou!")
				pc++
			}

			fmt.Println()
		}

		fmt.Println("PLACAR FINAL:")
		fmt.Printf("Você: %d | PC: %d\n", voce, pc)

		fmt.Println("JOGAR NOVAMENTE?")
		fmt.Println("1 - Sim")
		fmt.Println("0 - Sair")

		var resp int
		fmt.Scan(&resp)

		if resp == 0 {
			break
		}
	}
}