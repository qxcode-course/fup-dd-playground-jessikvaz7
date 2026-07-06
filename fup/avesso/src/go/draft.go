package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, x int
		fmt.Scan(&n, &x)

		v := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&v[i])
		}

		pos := -1
		for i := 0; i < n; i++ {
			val := v[i]
			if val < 0 {
				val = -val
			}
			if val == x {
				pos = i
				break
			}
		}

		if pos != -1 {
			if pos > 0 {
				v[pos-1] = -v[pos-1]
			}
			if pos < n-1 {
				v[pos+1] = -v[pos+1]
			}
		}

		fmt.Print("[")
		for i := 0; i < n; i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v[i])
		}
		fmt.Println("]")
	}
}