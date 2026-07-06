package main

import "fmt"

func main() {
	var v []int

	for {
		var x int
		n, err := fmt.Scan(&x)
		if n == 0 || err != nil {
			break
		}
		v = append(v, x)
	}

	fmt.Print("[ ")
	for i := len(v) - 1; i >= 0; i-- {
		fmt.Print(v[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println(" ]")
}