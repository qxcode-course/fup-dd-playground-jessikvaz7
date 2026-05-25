package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	dif := a - b

	fmt.Println(abs(dif))
}
