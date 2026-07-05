package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	delta := b*b - 4*a*c

	if delta > 0 {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)

		if math.Abs(x1) < 1e-9 {
			x1 = 0
		}
		if math.Abs(x2) < 1e-9 {
			x2 = 0
		}

		fmt.Printf("%.2f\n", x1)
		fmt.Printf("%.2f\n", x2)

	} else if delta == 0 {
		x := -b / (2 * a)

		if math.Abs(x) < 1e-9 {
			x = 0
		}

		fmt.Printf("%.2f\n", x)

	} else {
		fmt.Println("nao ha raiz real")
	}
}