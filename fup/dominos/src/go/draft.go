package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	ordenado := true

	for i := 1; i < n; i++ {
		if v[i] < v[i-1] {
			ordenado = false
			break
		}
	}

	if ordenado {
		fmt.Println("ok")
	} else {
		fmt.Println("precisa de ajuste")
	}
}