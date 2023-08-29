package main

import "fmt"

func main() {
	var numeros []int = make([]int, 0)

	for {
		var numero int
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		numeros = append(numeros, numero)
	}

	for _, numero := range numeros {
		fmt.Printf("f91(%d) = %d\n", numero, f91(numero))
	}
}

func f91(n int) int {
	if n <= 100 {
		return f91(f91(n + 11))
	}

	return n - 10
}