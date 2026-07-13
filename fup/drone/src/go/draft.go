package main

import "fmt"

func passa(x, y, h, l int) bool {
	return (x <= h && y <= l) || (x <= l && y <= h)
}

func main() {
	var a, b, c int
	var h, l int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&h)
	fmt.Scan(&l)

	if passa(a, b, h, l) ||
		passa(a, c, h, l) ||
		passa(b, c, h, l) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}