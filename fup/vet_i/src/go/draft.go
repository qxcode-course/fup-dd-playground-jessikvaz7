package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	numeros := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&numeros[i])
	}

	for i := 0; i < N; i++ {
		fmt.Println(numeros[i])
	}
}