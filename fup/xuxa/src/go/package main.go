package go
package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os,stdin)
	contador := make(map[rune]int)
	for scanner.Scan() {
	line := scanner.text()
    for _, c := range line {
	contador [c] +=1

	}
  }
  for char, cont := range line {
	fmt.print
}	