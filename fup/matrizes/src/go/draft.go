package main

import "fmt"

func main() {
	var l, c int
	fmt.Scan(&l)
	fmt.Scan(&c)

	A := make([][]int, l)
	B := make([][]int, l)

	for i := 0; i < l; i++ {
		A[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	for i := 0; i < l; i++ {
		B[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&B[i][j])
		}
	}

	for i := 0; i < l; i++ {
		fmt.Print("[ ")
		for j := 0; j < c; j++ {
			fmt.Print(A[i][j] + B[i][j])
			if j < c-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println(" ]")
	}
}