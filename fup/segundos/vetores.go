package main

import "fmt"

func main() {
	var tempo int

	fmt.Scan(&tempo)

	horas := tempo / 3600
	resto := tempo % 3600
	minutos := resto / 60
	segundos := resto % 60

	fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}