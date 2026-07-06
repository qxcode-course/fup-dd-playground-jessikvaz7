package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	soma := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		soma += x
	}

	media := float64(soma) / float64(n)

	fmt.Printf("%.1f\n", media)
}