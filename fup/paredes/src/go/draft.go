package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	maior := -1
	cont := 0

	for i := 0; i < n; i++ {
		var h int
		fmt.Scan(&h)

		if h > maior {
			cont++
			maior = h
		}
	}

	fmt.Println(cont)
}