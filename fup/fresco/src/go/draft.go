package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func vogal(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	texto, _ := reader.ReadString('\n')
	texto = strings.TrimSpace(texto)

	palavras := strings.Split(texto, " ")

	resp := palavras[0]

	for i := 1; i < len(palavras); i++ {
		ult := resp[len(resp)-1]
		pri := palavras[i][0]

		if vogal(ult) && vogal(pri) {
			// Remove apenas a primeira vogal da próxima palavra
			resp += palavras[i][1:]
		} else {
			resp += " " + palavras[i]
		}
	}

	fmt.Println(resp)
}