package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	mat := make([][]byte, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		mat[i] = []byte(s)
	}

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == '*' {
				continue
			}

			cont := 0
			for k := 0; k < 8; k++ {
				ni := i + dx[k]
				nj := j + dy[k]

				if ni >= 0 && ni < n && nj >= 0 && nj < m && mat[ni][nj] == '*' {
					cont++
				}
			}

			if cont > 0 {
				mat[i][j] = byte(cont + '0')
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(string(mat[i]))
	}
}