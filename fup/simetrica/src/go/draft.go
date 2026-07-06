package main

import "fmt"

func main() {
	var mat [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			if mat[i][j] != mat[j][i] {
				fmt.Println("nao")
				return
			}
		}
	}

	fmt.Println("sim")
}