package main

import "fmt"

func main() {
	var M, A, B int

	fmt.Scan(&M)
	fmt.Scan(&A)
	fmt.Scan(&B)

	C := M - A - B

	maior := A

	if B > maior {
		maior = B
	}

	if C > maior {
		maior = C
	}

	fmt.Println(maior)
}