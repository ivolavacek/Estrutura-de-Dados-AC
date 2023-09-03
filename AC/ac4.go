package main

import (
	"fmt"
	"math"
)

func main () {
	var pecas float64
	fmt.Print("Informe o número de peças: ")
	fmt.Scan(&pecas)

	hanoi(pecas)
}

func hanoi (n float64) {
	var minMovimentos float64

	minMovimentos = math.Pow(2, n) - 1

	fmt.Println(minMovimentos)
}