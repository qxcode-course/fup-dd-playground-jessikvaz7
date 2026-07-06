package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var chico, cebolinha, n int
	fmt.Scan(&chico)
	fmt.Scan(&cebolinha)
	fmt.Scan(&n)

	total := 0

	for i := 0; i < n; i++ {
		var animal string
		fmt.Scan(&animal)

		if animal == "g" {
			total += 2
		} else if animal == "v" || animal == "c" {
			total += 4
		}
	}

	fmt.Println(total)

	difChico := abs(chico - total)
	difCebolinha := abs(cebolinha - total)

	if difChico < difCebolinha {
		fmt.Println("Chico Bento")
	} else if difCebolinha < difChico {
		fmt.Println("Cebolinha")
	} else {
		fmt.Println("empate")
	}
}