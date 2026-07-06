package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mat := make([][]string, n)

	leaoL, leaoC := -1, -1

	for i := 0; i < n; i++ {
		mat[i] = make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&mat[i][j])
			if mat[i][j] == "L" {
				leaoL = i
				leaoC = j
			}
		}
	}

	gladiadores := 0
	condenados := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if leaoL != -1 && (i == leaoL || j == leaoC) {
				continue
			}

			if mat[i][j] == "G" {
				gladiadores += 2
			} else if mat[i][j] == "C" {
				if i+j == n-1 {
					condenados += 2
				} else {
					condenados++
				}
			}
		}
	}

	if gladiadores > condenados {
		fmt.Println("Gladiadores")
	} else if condenados > gladiadores {
		fmt.Println("Condenados a morte")
	} else {
		fmt.Println("Ninguem")
	}
}