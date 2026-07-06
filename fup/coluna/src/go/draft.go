package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	maiorColuna := 0
	maiorValor := -1

	for j := 0; j < n; j++ {
		soma := 0
		for i := 0; i < n; i++ {
			soma += mat[i][j] * mat[i][j]
		}

		if soma > maiorValor {
			maiorValor = soma
			maiorColuna = j
		}
	}

	fmt.Println(maiorColuna)
}