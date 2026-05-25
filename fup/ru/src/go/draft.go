package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	frase, _ := reader.ReadString('\n')

	vogais := ""
	consoantes := ""

	for _, letra := range frase {

		if letra == ' ' || letra == '\n' {
			continue
		}

		if letra == 'a' ||
			letra == 'e' ||
			letra == 'i' ||
			letra == 'o' ||
			letra == 'u' {

			vogais += string(letra)

		} else {

			consoantes += string(letra)
		}
	}

	fmt.Println(vogais)
	fmt.Println(consoantes)
}