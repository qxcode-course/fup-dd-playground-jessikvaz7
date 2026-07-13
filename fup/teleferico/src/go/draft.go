package main

import "fmt"

func main() {
	var c, a int

	fmt.Scan(&c)
	fmt.Scan(&a)

	viagens := (a + (c - 2)) / (c - 1)

	fmt.Println(viagens)
}