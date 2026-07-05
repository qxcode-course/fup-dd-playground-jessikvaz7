package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Print("[")
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(vetor[i])
	}
	fmt.Println("]")
}