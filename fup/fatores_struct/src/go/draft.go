package main

import "fmt"

type Fator struct {
	num int
	qtd int
}

func calcFatores(n int) []Fator {
	var fatores []Fator

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			cont := 0
			for n%i == 0 {
				n /= i
				cont++
			}
			fatores = append(fatores, Fator{i, cont})
		}
	}

	if n > 1 {
		fatores = append(fatores, Fator{n, 1})
	}

	return fatores
}

func main() {
	var n int
	fmt.Scan(&n)

	fatores := calcFatores(n)

	for _, f := range fatores {
		fmt.Println(f.num, f.qtd)
	}
}