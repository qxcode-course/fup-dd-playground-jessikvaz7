package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	var i, q int

	fmt.Fscanln(reader, &i)
	fmt.Fscanln(reader, &q)

	if i < 0 || i >= len(s) {
		fmt.Println("")
		return
	}

	if q <= 0 {
		fmt.Println("")
		return
	}

	fim := i + q

	if fim > len(s) {
		fim = len(s)
	}

	fmt.Println(s[i:fim])
}