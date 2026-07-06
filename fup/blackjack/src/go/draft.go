package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	soma := 0
	ases := 0

	for i := 0; i < n; i++ {
		var carta int
		fmt.Scan(&carta)

		if carta == 1 {
			soma += 11
			ases++
		} else if carta >= 11 && carta <= 13 {
			soma += 10
		} else {
			soma += carta
		}
	}

	for soma > 21 && ases > 0 {
		soma -= 10
		ases--
	}

	fmt.Println(soma)
}