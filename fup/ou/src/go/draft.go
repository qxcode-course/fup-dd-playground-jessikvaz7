package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
 if num == 3 || num == 5 {
	fmt.Println("SIM")
 } else {
	fmt.Println("NAO")
  }
}