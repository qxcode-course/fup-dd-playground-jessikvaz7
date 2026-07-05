package main

import (
	"fmt"
	"sort"
)

type Item struct {
	valor  int
	indice int
}

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]Item, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i].valor)
		v[i].indice = i
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i].valor < v[j].valor
	})

	fmt.Print("[ ")
	for i := 0; i < n; i++ {
		fmt.Print(v[i].indice)
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println(" ]")
}