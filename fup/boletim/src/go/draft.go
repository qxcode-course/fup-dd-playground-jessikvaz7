package main

import "fmt"

func main() {
	var mat [2][3]int
	soma := 0

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&mat[i][j])
			soma += mat[i][j]
		}
	}

	fmt.Println(soma)
}