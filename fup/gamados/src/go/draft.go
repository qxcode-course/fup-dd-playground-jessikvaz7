package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)

	palavras := strings.Split(frase, " ")

	ordenado := true

	for i := 1; i < len(palavras); i++ {
		if palavras[i-1] > palavras[i] {
			ordenado = false
			break
		}
	}

	if ordenado {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}
}