package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	mat := make([][]int, n)
	colunas := make([]int, m)

	maior := 0

	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
		somaLinha := 0

		for j := 0; j < m; j++ {
			fmt.Scan(&mat[i][j])
			somaLinha += mat[i][j]
			colunas[j] += mat[i][j]
		}

		if somaLinha > maior {
			maior = somaLinha
		}
	}

	for j := 0; j < m; j++ {
		if colunas[j] > maior {
			maior = colunas[j]
		}
	}

	fmt.Println(maior)
}