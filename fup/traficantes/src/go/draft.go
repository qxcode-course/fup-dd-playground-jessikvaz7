package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	texto, _ := reader.ReadString('\n')
	texto = strings.TrimRight(texto, "\r\n")

	antiga, _ := reader.ReadString('\n')
	antiga = strings.TrimRight(antiga, "\r\n")

	nova, _ := reader.ReadString('\n')
	nova = strings.TrimRight(nova, "\r\n")

	i := 0
	for i < len(texto) {
		if i+len(antiga) <= len(texto) && texto[i:i+len(antiga)] == antiga {
			fmt.Print(nova)
			i += len(antiga)
		} else {
			fmt.Printf("%c", texto[i])
			i++
		}
	}
	fmt.Println()
}