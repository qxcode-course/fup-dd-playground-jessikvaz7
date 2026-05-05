package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	soma := a + b + c

	fmt.Println(soma)
}