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

    for i := len(frase) -1; i >= 0; i-- {
        fmt.Print(string(frase[i]))
    }
    fmt.Println()
}
