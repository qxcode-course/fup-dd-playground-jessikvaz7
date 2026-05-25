package main

import (
	"fmt"
	"strconv"
)

func filtrarimpares(nums []int) []int {
	impares := []int{}

	for _, elem := range nums {
		if elem%2 == 1 {
			impares = append(impares, elem)
		}
	}

	return impares
}

func index(nums []int, valor int) int {
	for i, elem := range nums {
		if elem == valor {
			return i
		}
	}

	return -1
}

func contains(nums []int, valor int) bool {
	for _, elem := range nums {
		if elem == valor {
			return true
		}
	}

	return false
}

func count(nums []int, valor int) int {
	contador := 0

	for _, elem := range nums {
		if elem == valor {
			contador++
		}
	}

	return contador
}

func separarfigurinhas(montante []int) ([]int, []int) {
	album := []int{}
	repet := []int{}

	for _, fig := range montante {
		if !contains(album, fig) {
			album = append(album, fig)
		} else {
			repet = append(repet, fig)
		}
	}

	return album, repet
}

func main() {
	num, err := strconv.Atoi("3641")

	if err == nil {
		fmt.Println(num)
	}
}