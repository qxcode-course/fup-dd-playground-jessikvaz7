package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	vacinas := make([]int, n)
	pacientes := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vacinas[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&pacientes[i])
	}

	sort.Ints(vacinas)
	sort.Ints(pacientes)

	for i := 0; i < n; i++ {
		if vacinas[i] <= pacientes[i] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}