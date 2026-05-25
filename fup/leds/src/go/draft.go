package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	leds := map[rune]int{
		'0': 6,
		'1': 2,
		'2': 5,
		'3': 5,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 3,
		'8': 7,
		'9': 6,
	}

	for i := 0; i < n; i++ {
		var valor string

		fmt.Scan(&valor)

		total := 0

		for _, digito := range valor {
			total += leds[digito]
		}

		fmt.Printf("%d leds\n", total)
	}
}