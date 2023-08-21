package main

import "fmt"

func main() {
	// Array é uma coleção de dados do mesmo tipo, com tamanho definido em compilação
	var filmes [5]string

	filmes[0] = "O Senhor dos Aneis"
	filmes[1] = "O Poderoso Chefão"
	fmt.Println(filmes)

	// Delaração curta
	numeros := [4]int{0, 2, 4, 6}

	fmt.Println(numeros)

	/* Slices são segmentos de arrays, que podem ou não ter sido definidos
	previamente. Eles possuem dimensão dinâmica. */
	var outrosNumeros []int

	outrosNumeros = numeros[1:]  // Elemento 1 até o final
	outrosNumeros = numeros[1:3] // Elemento 1 (inclusive) até 3 (exclusive)

	fmt.Println(outrosNumeros)

	// Não precisa criar um slice a partir de um array previamente definido
	nomes := []string{"Ana", "Pedro"}

	fmt.Println(nomes)

	// Adicionar elementos no slice
	nomes = append(nomes, "João")
	fmt.Println(nomes)

	nomes = append(nomes[:1], nomes[2])
	fmt.Println(nomes)

	// Iterando sobre arrays e slices
	fmt.Println("----------------")
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	for indice, num := range numeros {
		fmt.Println("Índice", indice, "- valor", num)
	}

	fmt.Println("----------------")
	modificaArray(numeros) // Passagem de valor dos arrays
	fmt.Println(numeros)

	outrosNumeros = numeros[1:3]
	modificaSlice(outrosNumeros) // Passagem por referência dos slices
	fmt.Println(outrosNumeros)
	fmt.Println(numeros)
}


func modificaArray(a [4]int) {
	a[0] = 999
	fmt.Println(a)
}

func modificaSlice(s []int) {
	s[0] = 999
}

func criarSlice() []int {
	novoSlice := []int{10, 20, 30}
	return novoSlice
}