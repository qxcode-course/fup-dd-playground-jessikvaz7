package main

import "fmt"

func main() {
	var N, Q int
	fmt.Scan(&N, &Q)

	agua := make([]int, N)

	for i := 0; i < Q; i++ {
		var A, B, L int
		fmt.Scan(&A, &B, &L)

		for j := A; j <= B; j++ {
			agua[j] += L
		}
	}

	for i := 0; i < N; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(agua[i])
	}
	fmt.Println()
}