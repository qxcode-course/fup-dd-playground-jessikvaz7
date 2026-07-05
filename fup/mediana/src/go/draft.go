package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	sort.Float64s(v)

	var mediana float64

	if n%2 == 1 {
		mediana = v[n/2]
	} else {
		mediana = (v[n/2-1] + v[n/2]) / 2
	}

	fmt.Printf("%.1f\n", mediana)
}