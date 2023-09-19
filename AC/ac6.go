package main

import "fmt"

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	fmt.Println("Tentando inserir:", novoValor)
	if *n < M {
		if buscaBin(*v, *n, novoValor) == -1 { // confere se 'novoValor' ja existe na lista
			fmt.Println(novoValor, "não encontrado na lista, inserindo agora")
			i := *n - 1
			for i >= 0 {
				if v[i] > novoValor { // para todo valor maior que o 'novoValor', coloca um índice a direita, sabendo que não há valores repetidos
					v[i + 1] = v[i]
				} else {
					break
				}
				i-- // diminui o índice para percorrer a lista da direta para esquerda
			}
			v[i + 1] = novoValor // adiciona o 'novoValor' no primeiro espaço em que 'novoValor' < v[i]
			*n++
		} else {
			fmt.Println(novoValor, "já existe na lista")
		}
	} else {
		fmt.Println("Overflow!")
	}
}

func buscaBin(v [M]int, n, x int) int {
	inf, sup := 0, n - 1
	for inf <= sup {
		meio := int((inf + sup) / 2)
		if v[meio] == x {
			return meio
		}
		if v[meio] < x {
			inf = meio + 1
		} else {
			sup = meio - 1
		}
	}
	return -1
}