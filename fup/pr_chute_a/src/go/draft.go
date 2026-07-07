package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	inferior := 0
	superior := 100
	secreto := rand.Intn(99) + 1

	for {
		// Verifica se só resta um número possível
		if superior-inferior == 2 {
			fmt.Printf("Era %d, você perdeu!\n", secreto)
			break
		}

		fmt.Printf("Diga um numero entre ]%d, %d[: ", inferior, superior)

		var chute int
		fmt.Scan(&chute)

		// Garante que o chute esteja dentro do intervalo aberto
		if chute <= inferior || chute >= superior {
			fmt.Println("Chute invalido!")
			continue
		}

		if chute == secreto {
			fmt.Printf("Era %d, você ganhou!\n", secreto)
			break
		}

		if chute < secreto {
			inferior = chute
		} else {
			superior = chute
		}
	}
}