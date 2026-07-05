package main

import "fmt"

func main() {
	var h1, m1, s1 int
	var h2, m2, s2 int

	fmt.Scan(&h1, &m1, &s1)
	fmt.Scan(&h2, &m2, &s2)

	inicio := h1*3600 + m1*60 + s1
	fim := h2*3600 + m2*60 + s2

	if fim < inicio {
		fim += 24 * 3600
	}

	diferenca := fim - inicio

	horas := diferenca / 3600
	diferenca = diferenca % 3600

	minutos := diferenca / 60
	segundos := diferenca % 60

	fmt.Printf("%02d %02d %02d\n", horas, minutos, segundos)
}