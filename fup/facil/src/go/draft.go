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

	if ida < volta {
		fmt.Println(ida)
	} else {
		fmt.Println(volta)
	}
}