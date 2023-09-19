package main

import "fmt"

/*
PROGRAMA PUSH(v, topo, M, valor)
	Se topo = M - 1 então:
		"Overflow"
	Senão:
		topo++
		v[topo] = valor

PROGRAMA POP(v, topo)
	Se topo = -1 então
		"Underflow"
		Retorna -1
	Senão
		valorRemovido := v[topo]
		Limpar v[topo]
		topo--
		Retorna valorRemovido
*/

const M = 3

func main() {
	var pilha [M]int
	topo := -1

	push(&pilha, &topo, 4)
	fmt.Println(pilha)
	push(&pilha, &topo, 5)
	fmt.Println(pilha)
	push(&pilha, &topo, 2)
	fmt.Println(pilha)
	push(&pilha, &topo, 3) // Overflow
	fmt.Println(pilha)

	fmt.Println(pop(&pilha, &topo))
	fmt.Println(pilha)
	fmt.Println(pop(&pilha, &topo))
	fmt.Println(pilha)
	fmt.Println(pop(&pilha, &topo))
	fmt.Println(pilha)
	fmt.Println(pop(&pilha, &topo)) // Underflow
	fmt.Println(pilha)
}

func pop(v *[M]int, topo *int) int {
	if *topo == -1 {
		fmt.Println("Underflow")
		return -1
	}

	valorRemovido := v[*topo]
	v[*topo] = 0
	*topo--
	return valorRemovido
}

func push(v *[M]int, topo *int, valor int) {
	if *topo == M - 1 {
		fmt.Println("Overflow")
		return
	}

	*topo++
	v[*topo] = valor
}