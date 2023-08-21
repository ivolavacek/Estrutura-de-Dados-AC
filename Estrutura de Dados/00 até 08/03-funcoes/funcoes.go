package main

import "fmt"

func main() {
	fmt.Println(soma(2, 6))
	fmt.Println(subtracao(4, 7))
	fmt.Println(trocaValores("oi", "olá"))

	usoAnonima()
}

//funções anônimas (closure)
func usoAnonima() {
	dobra := func(x int) int {
		return x * 2
	}

	resultado := dobra(5)
	fmt.Println(resultado)

	calcular := func(f func(int) int, x int) int {
		return f(x)
	}

	resultado = calcular(dobra, 4)
	fmt.Println(resultado)
}

// func nome(param1 tipo, param2 tipo, ...) tipoRetorno {}
func soma(a int, b int) int {
	return a + b
}

//parâmetros com tipo repetido podem ser omitidos
func subtracao(a, b float64) float64 {
	return a - b
}

// retornos múltiplos
func trocaValores(a, b string) (string, string) {
	return b, a
}
