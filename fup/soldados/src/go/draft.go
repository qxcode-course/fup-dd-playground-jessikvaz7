package main

import "fmt"

func main() {
	var nl, nc int
	fmt.Scan(&nl, &nc)

	mat := make([][]int, nl)
	for i := 0; i < nl; i++ {
		mat[i] = make([]int, nc)
		for j := 0; j < nc; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	cont := 0

	for j := 0; j < nc; j++ {
		for i := 0; i < nl-1; i++ {
			if mat[i+1][j] < mat[i][j] {
				cont++
			}
		}
	}

	fmt.Println(cont)
}