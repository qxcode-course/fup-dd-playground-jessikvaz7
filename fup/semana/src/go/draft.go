package main

import "fmt"

func main() {
	var dia, hora int

	fmt.Scan(&dia)
	fmt.Scan(&hora)

	if dia >= 2 && dia <= 6 { // Segunda a sexta
		if (hora >= 8 && hora <= 11) || (hora >= 14 && hora <= 17) {
			fmt.Println("SIM")
		} else {
			fmt.Println("NAO")
		}
	} else if dia == 7 { // Sábado
		if hora >= 8 && hora <= 11 {
			fmt.Println("SIM")
		} else {
			fmt.Println("NAO")
		}
	} else { // Domingo
		fmt.Println("NAO")
	}
}