package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	num := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		num = num*10 + x
	}

	fmt.Println(num)
}