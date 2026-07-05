package main

import "fmt"

func main() {
	var B, T int

	fmt.Scan(&B)
	fmt.Scan(&T)

	area := float64(B+T) * 70 / 2

	if area > 5600 {
		fmt.Println(1)
	} else if area < 5600 {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}