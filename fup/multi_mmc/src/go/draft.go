package main

import "fmt"

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmc(v []int) int {
	m := v[0]

	for i := 1; i < len(v); i++ {
		m = (m * v[i]) / mdc(m, v[i])
	}

	return m
}

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println(mmc(v))
}