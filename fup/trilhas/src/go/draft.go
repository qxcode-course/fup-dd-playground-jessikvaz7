package main

import "fmt"

func esforco(v []int) int {
	soma := 0
	for i := 1; i < len(v); i++ {
		if v[i] > v[i-1] {
			soma += v[i] - v[i-1]
		}
	}
	return soma
}

func main() {
	var n int
	fmt.Scan(&n)

	melhorTrilha := 1
	menorEsforco := -1

	for t := 1; t <= n; t++ {
		var m int
		fmt.Scan(&m)

		h := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&h[i])
		}

		ida := esforco(h)

		inv := make([]int, m)
		for i := 0; i < m; i++ {
			inv[i] = h[m-1-i]
		}

		volta := esforco(inv)

		e := ida
		if volta < e {
			e = volta
		}

		if menorEsforco == -1 || e < menorEsforco {
			menorEsforco = e
			melhorTrilha = t
		}
	}

	fmt.Println(melhorTrilha)
}