package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	vencedor := -1
	menor := 101

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		if a >= 10 && b >= 10 {
			dif := abs(a - b)

			if vencedor == -1 || dif < menor {
				menor = dif
				vencedor = i
			}
		}
	}

	if vencedor == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(vencedor)
	}
}