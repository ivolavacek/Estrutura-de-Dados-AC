package main

import (
	"fmt"
	"strconv"
)

/*
9 -> grau-9 1
18 -> grau-9 1
189 -> grau-9 2
999999999999999999999 -> grau-9 3

99999...9999 -> 1000 dígitos

PROGRAMA PRINCIPAL
	Lê numero
	Escreve numero + RUMO_AOS_9(numero, 0)

PROGRAMA RUMO_AOS_9(n, grau)
	Se n = "9" então:
		Se grau = 0 então grau = 1
		Retorna " is a multiple of 9 and has 9-degree " + grau + "."
	Senão se o comprimento de n for igual a 1 então:
		Retorna " is not a multiple of 9."
	Senão:
		soma := 0

		Para cada digito em n, faça:
			Converter digito para inteiro
			soma += digito

		Converter soma para string
		Retorna RUMO_AOS_9(soma, grau + 1)

n -> "9":
	"9 is a multiple of 9 and has 9-degree 1."
n -> "16":
	soma -> 7
	RUMO_AOS_9("7", 1):
		"16 is not a multiple of 9."
n -> "18":
	soma -> 9
	RUMO_AOS_9("9", 2):
		"18 is a multiple of 9 and has 9-degree 1."
n -> "189":
	soma -> 18
	RUMO_AOS_9("18", 1):
		soma -> 9
		RUMO_AOS_9("9", 2):
			"189 is a multiple of 9 and has 9-degree 2."
n -> "999999999999999999999":
	soma -> 189
	RUMO_AOS_9("189", 1):
		soma -> 18
		RUMO_AOS_9("18", 2):
			soma -> 9
			RUMO_AOS_9("9", 3):
				"999999999999999999999 is a multiple of 9 and has 9-degree 3."
*/
func main() {
	var numeros []string = make([]string, 0)

	for {
		var strNumero string
		fmt.Scanln(&strNumero)

		if strNumero == "0" {
			break
		}

		numeros = append(numeros, strNumero)
	}

	for _, strNumero := range numeros {
		fmt.Println(strNumero + rumoAos9(strNumero, 0))
	}
}

func rumoAos9(n string, grau int) string {
	if n == "9" {
		if grau == 0 {
			grau = 1
		}

		return " is a multiple of 9 and has 9-degree " + strconv.Itoa(grau) + "."
	}

	if len(n) == 1 {
		return " is not a multiple of 9."
	}

	soma := 0
	for _, digito := range n {
		soma += int(digito - '0') // convertendo rune de um caractere unicode numérico para um inteiro
	}

	return rumoAos9(strconv.Itoa(soma), grau + 1)
}