package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	texto, _ := reader.ReadString('\n')

	for _, c := range texto {
		if unicode.IsLower(c) {
			fmt.Printf("%c", unicode.ToUpper(c))
		} else if unicode.IsUpper(c) {
			fmt.Printf("%c", unicode.ToLower(c))
		} else {
			fmt.Printf("%c", c)
		}
	}
}