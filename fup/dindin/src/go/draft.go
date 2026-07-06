package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	chocolate := 0
	limao := 0
	manha := 0
	tarde := 0

	for i := 0; i < n; i++ {
		var sabor, turno string
		fmt.Scan(&sabor, &turno)

		if sabor == "c" {
			chocolate++
		} else if sabor == "l" {
			limao++
		}

		if turno == "m" {
			manha++
		} else if turno == "t" {
			tarde++
		}
	}

	if chocolate > limao {
		fmt.Println("c")
	} else if limao > chocolate {
		fmt.Println("l")
	} else {
		fmt.Println("empate")
	}

	if manha < tarde {
		fmt.Println("m")
	} else if tarde < manha {
		fmt.Println("t")
	} else {
		fmt.Println("empate")
	}
}