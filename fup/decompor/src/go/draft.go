package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var v []int

	if n == 0 {
		fmt.Println(0)
		return
	}

	for n > 0 {
		v = append(v, n%10)
		n /= 10
	}

	for i := len(v) - 1; i >= 0; i-- {
		fmt.Print(v[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}