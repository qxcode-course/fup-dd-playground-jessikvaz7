package main

import "fmt"

func main() {
	var m [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&m[i][j])
		}
	}

	soma := m[0][0] + m[0][1] + m[0][2]

	for i := 0; i < 3; i++ {
		linha := 0
		coluna := 0

		for j := 0; j < 3; j++ {
			linha += m[i][j]
			coluna += m[j][i]
		}

		if linha != soma || coluna != soma {
			fmt.Println("nao")
			return
		}
	}

	diag1 := m[0][0] + m[1][1] + m[2][2]
	diag2 := m[0][2] + m[1][1] + m[2][0]

	if diag1 != soma || diag2 != soma {
		fmt.Println("nao")
	} else {
		fmt.Println("sim")
	}
}