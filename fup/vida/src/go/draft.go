package main

import "fmt"

func main() {
	var nl, nc int
	fmt.Scan(&nl, &nc)

	mat := make([][]byte, nl)
	nova := make([][]byte, nl)

	for i := 0; i < nl; i++ {
		var linha string
		fmt.Scan(&linha)
		mat[i] = []byte(linha)
		nova[i] = make([]byte, nc)
	}

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			vivos := 0

			for k := 0; k < 8; k++ {
				ni := i + dx[k]
				nj := j + dy[k]

				if ni >= 0 && ni < nl && nj >= 0 && nj < nc {
					if mat[ni][nj] == '#' {
						vivos++
					}
				}
			}

			if mat[i][j] == '#' {
				if vivos == 2 || vivos == 3 {
					nova[i][j] = '#'
				} else {
					nova[i][j] = '.'
				}
			} else {
				if vivos == 3 {
					nova[i][j] = '#'
				} else {
					nova[i][j] = '.'
				}
			}
		}
	}

	for i := 0; i < nl; i++ {
		fmt.Println(string(nova[i]))
	}
}