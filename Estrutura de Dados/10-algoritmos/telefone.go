package main

import "fmt"

/*
PROGRAMA ENCONTRE_TELEFONE
	Lê expressoes

	Para cada expressao em expressoes, faça:
		telefone := ""

		Para cada caractere em expressao, faça:
			Analise caractere:
				Caso "1", "0" ou "-":
					telefone += caractere
				Caso "A", "B" ou "C":
					telefone += "2"
				Caso "D", "E" ou "F":
					telefone += "3"
				Caso "G", "H" ou "I":
					telefone += "4"
				Caso "J", "K" ou "L":
					telefone += "5"
				Caso "M", "N" ou "O":
					telefone += "6"
				Caso "P", "Q", "R" ou "S":
					telefone += "7"
				Caso "T", "U" ou "V":
					telefone += "8"
				Caso "W", "X", "Y" ou "Z":
					telefone += "9"

		Escreve telefone
*/
func main() {
	var expressoes []string = make([]string, 0)

	for {
		var expressao string
		fmt.Scanln(&expressao)

		if expressao == "" {
			break
		}

		expressoes = append(expressoes, expressao)
	}

	for _, expressao := range expressoes {
		telefone := ""

		for _, caractere := range expressao {
			switch caractere {
			case '1', '0', '-':
				telefone += string(caractere)
			case 'A', 'B', 'C':
				telefone += "2"
			case 'D', 'E', 'F':
				telefone += "3"
			case 'G', 'H', 'I':
				telefone += "4"
			case 'J', 'K', 'L':
				telefone += "5"
			case 'M', 'N', 'O':
				telefone += "6"
			case 'P', 'Q', 'R', 'S':
				telefone += "7"
			case 'T', 'U', 'V':
				telefone += "8"
			case 'W', 'X', 'Y', 'Z':
				telefone += "9"
			}
		}

		fmt.Println(telefone)
	}
}