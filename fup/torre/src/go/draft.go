package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mat := make([][]int, n)
	somaLinhas := make([]int, n)
	somaColunas := make([]int, n)

	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&mat[i][j])
			somaLinhas[i] += mat[i][j]
			somaColunas[j] += mat[i][j]
		}
	}

	maior := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			peso := somaLinhas[i] + somaColunas[j] - 2*mat[i][j]
			if peso > maior {
				maior = peso
			}
		}
	}

	fmt.Println(maior)
}