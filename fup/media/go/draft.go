package main
import "fmt"
func main() {
	var num1 float64
	var num2 float64
	var media float64
	fmt.Scan(&num1, &num2)
	media = (num1 + num2) / 2.0
	fmt.Printf("%.1f\n", media)

    
}
