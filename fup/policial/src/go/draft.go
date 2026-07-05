package main

import "fmt"

func bubbleSort(v []int) {
	n := len(v)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	bubbleSort(v)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v[i])
	}
	fmt.Println()
}