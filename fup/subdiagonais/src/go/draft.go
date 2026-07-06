package main

import "fmt"

func main() {
	var mat [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	somaPrincipal := 0
	somaSecundaria := 0

	for i := 0; i < 5; i++ {
		somaPrincipal += mat[i][i]
		somaSecundaria += mat[i][4-i]
	}

	fmt.Println(somaPrincipal - somaSecundaria)
}