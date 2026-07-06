package main

import "fmt"

func main() {
	teste := 1
	primeiro := true

	for {
		var p, r int
		fmt.Scan(&p, &r)

		if p == 0 && r == 0 {
			break
		}

		if !primeiro {
			fmt.Println()
		}
		primeiro = false

		fila := make([]int, p)
		for i := 0; i < p; i++ {
			fmt.Scan(&fila[i])
		}

		for k := 0; k < r; k++ {
			var n, j int
			fmt.Scan(&n, &j)

			nova := make([]int, 0)

			for i := 0; i < n; i++ {
				var a int
				fmt.Scan(&a)

				if a == j {
					nova = append(nova, fila[i])
				}
			}

			fila = nova
		}

		fmt.Printf("Teste %d\n", teste)
		fmt.Println(fila[0])

		teste++
	}
}