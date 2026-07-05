package main

import "fmt"

func imprimir(v []int) {
	fmt.Print("[")

	for i := 0; i < len(v); i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v[i])
	}

	fmt.Println("]")
}

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	imprimir(v)
}