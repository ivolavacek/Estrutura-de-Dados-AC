package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
PROGRAMA CARTAS
	n1, n2 := 0
	classificacao := ""

	Lê n1
	Para cada i entre 2 e 5, faça:
		Lê n2

		Se n2 > n1 então:
			Se classificacao = "" então:
				classificacao = "C"
			Senão se classificacao = "D" então:
				classificacao = "N"
				Pare

		Se n1 > n2 então:
			Se classificacao = "" então:
				classificacao = "D"
			Senão se classificacao = "C" então:
				classificacao = "N"
				Pare

		n1 = n2

	Escreve classificacao
*/
func main() {
	n1, n2 := 0, 0
	classificacao := ""

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Lê as informações do buffer

	linha := scanner.Text() // Converte o buffer para string
	numeros := strings.Split(linha, " ") // Divide a linha e vários números, no formato de string

	//{"1", "2", "3", "5", "6"}

	n1, _ = strconv.Atoi(numeros[0])
	for i := 1; i <= 4; i++ {
		// Retorna o valor convertido e uma mensagem de erro
		n2, _ = strconv.Atoi(numeros[i])

		if n2 > n1 {
			if classificacao == "" {
				classificacao = "C"
			} else if classificacao == "D" {
				classificacao = "N"
				break
			}
		} else {
			if classificacao == "" {
				classificacao = "D"
			} else if classificacao == "C" {
				classificacao = "N"
				break
			}
		}

		n1 = n2
	}

	fmt.Println(classificacao)
}