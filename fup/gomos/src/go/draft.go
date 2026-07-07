package main

import "fmt"

type Gomo struct {
	x, y int
}

func main() {
	var q int
	var d string

	fmt.Scan(&q, &d)

	cobra := make([]Gomo, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&cobra[i].x, &cobra[i].y)
	}

	// Guarda a posição antiga da cabeça
	antX, antY := cobra[0].x, cobra[0].y

	// Move a cabeça
	switch d {
	case "L":
		cobra[0].x--
	case "R":
		cobra[0].x++
	case "U":
		cobra[0].y--
	case "D":
		cobra[0].y++
	}

	// Cada gomo ocupa a posição anterior do da frente
	for i := 1; i < q; i++ {
		tmpX, tmpY := cobra[i].x, cobra[i].y
		cobra[i].x = antX
		cobra[i].y = antY
		antX, antY = tmpX, tmpY
	}

	for i := 0; i < q; i++ {
		fmt.Println(cobra[i].x, cobra[i].y)
	}
}