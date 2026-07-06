package main

import "fmt"

func main() {
	bingo := [4][4]int{
		{1, 9, 27, 23},
		{34, 20, 37, 47},
		{30, 87, 55, 69},
		{13, 60, 99, 66},
	}

	acertos := 0

	for i := 0; i < 6; i++ {
		var x int
		fmt.Scan(&x)

		for l := 0; l < 4; l++ {
			for c := 0; c < 4; c++ {
				if bingo[l][c] == x {
					acertos++
				}
			}
		}
	}

	fmt.Println(acertos)
}