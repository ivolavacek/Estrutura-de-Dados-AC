package main

import "fmt"

/*
PROGRAMA INSERE(v, in, fim, M, valor)
	Se (fim + 1) % M = in então:
		"Overflow"
	Senão:
		fim++
		Se fim = M então:
			fim = 0
		v[fim] = valor
		Se in = -1 então:
			in = 0

PROGRAMA REMOVE(v, in, fim)
	Se in = fim = -1 então:
		"Underflow"
	Senão:
		Limpar v[in]
		Se in = fim então:
			in = fim = -1
		Senão:
			in++
			Se in = M então:
				in = 0
*/

const M = 4

func main() {
	var fila [M]int
	in, fim := -1, -1

	insere(&fila, &in, &fim, 4)
	fmt.Println(fila)
	insere(&fila, &in, &fim, 5)
	fmt.Println(fila)
	insere(&fila, &in, &fim, 2)
	fmt.Println(fila)
	fmt.Println(remove(&fila, &in, &fim))
	fmt.Println(fila)
	insere(&fila, &in, &fim, 1)
	fmt.Println(fila)
	insere(&fila, &in, &fim, 7)
	fmt.Println(fila)

	fmt.Println(remove(&fila, &in, &fim))
	fmt.Println(fila)
	fmt.Println(remove(&fila, &in, &fim))
	fmt.Println(fila)
	fmt.Println(remove(&fila, &in, &fim))
	fmt.Println(fila)
	fmt.Println(remove(&fila, &in, &fim))
	fmt.Println(fila)
	fmt.Println(remove(&fila, &in, &fim)) // Underflow
	fmt.Println(fila)
}

func remove(f *[M]int, in *int, fim *int) int {
	if *in == *fim && *fim == -1 {
		fmt.Println("Underflow")
		return -1
	}

	valorRemovido := f[*in]
	f[*in] = 0

	if *in == *fim {
		*in, *fim = -1, -1
	} else {
		*in++
		if *in == M {
			*in = 0
		}
	}
	return valorRemovido
}

func insere(f *[M]int, in *int, fim *int, valor int) {
	if (*fim + 1) % M == *in {
		fmt.Println("Overflow")
		return
	}

	*fim++
	if *fim == M {
		*fim = 0
	}
	f[*fim] = valor

	if *in == -1 {
		*in = 0
	}
}