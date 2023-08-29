package main

import (
	"fmt"
	"strconv"
)

/*
PROGRAMA METEOROS
	contFazenda := 0
	mensagem := ""

	Enquanto VERDADEIRO, faça:
		Lê x1, y1, x2, y2
		Se x1 = y1 = x2 = y2 = 0 então Pare

		Lê numMeteoritos
		meteoritosNaFazenda := 0

		Para cada meteorito entre 1 e numMeteoritos, faça:
			Lê x, y
			Se x1 <= x <= 2 e y2 <= y <= y1 então:
				meteoritosNaFazenda++

		contFazenda++
		mensagem += "Teste " + contFazenda + meteoritosNaFazenda

	Escreve mensagem
*/
func main() {
	var x, y, x1, x2, y1, y2, numMeteoritos int
	contFazenda := 0
	mensagem := ""

	for {
		fmt.Scanln(&x1, &y1, &x2, &y2)
		if (x1 == 0 && x2 == 0 && y1 == 0 && y2 == 0) {
			break
		}

		fmt.Scanln(&numMeteoritos)
		meteoritosNaFazenda := 0

		for meteorito := 1; meteorito <= numMeteoritos; meteorito++ {
			fmt.Scanln(&x, &y)
			if x1 <= x && x <= x2 && y2 <= y && y <= y1 {
				meteoritosNaFazenda++
			}
		}

		contFazenda++
		mensagem += "Teste " + strconv.Itoa(contFazenda)
		mensagem += "\n" + strconv.Itoa(meteoritosNaFazenda) + "\n\n"
	}

	fmt.Print(mensagem)
}