package main

import "fmt"

/*
PROGRAMA MAIOR_NUMERO
	maior := 0

	Enquanto VERDADEIRO, faça:
		Lê numero
		Se numero = 0 então Pare
		Se numero > maior então:
			maior = numero

	Escreve maior
*/
func main() {
	var numero int
	maior := 0

	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}

		if numero > maior {
			maior = numero
		}
	}

	fmt.Println(maior)
}