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
	texto = strings.TrimSpace(texto)

	trecho, _ := reader.ReadString('\n')
	trecho = strings.TrimSpace(trecho)

	fmt.Println(strings.Count(texto, trecho))
}