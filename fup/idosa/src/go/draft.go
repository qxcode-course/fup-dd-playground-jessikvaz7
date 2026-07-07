package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
	sexo  string
}

func main() {
	var n int
	fmt.Scan(&n)

	pessoas := make([]Pessoa, n)

	maiorIdade := -1
	nome := ""

	for i := 0; i < n; i++ {
		fmt.Scan(&pessoas[i].nome, &pessoas[i].idade, &pessoas[i].sexo)

		if pessoas[i].sexo == "f" && pessoas[i].idade > maiorIdade {
			maiorIdade = pessoas[i].idade
			nome = pessoas[i].nome
		}
	}

	if maiorIdade == -1 {
		fmt.Println("nao tem mulher")
	} else {
		fmt.Println(nome)
	}
}