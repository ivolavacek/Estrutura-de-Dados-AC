package main

import "fmt"

func main() {
	var x = 10
	var dia = "segunda"

	if x > 18 {
		fmt.Println("Você é maior de idade")
	} else if x < 0 {
		fmt.Println("Valor inválido")
	} else {
		fmt.Println("Você é menor de idade")
	}

	switch dia {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia útil.")
	case "sábado", "domingo":
		fmt.Println("Final de semana.")
	default:
		fmt.Println("Dia inválido")
	}

	fmt.Println("-------------------")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------------")

	x = 5
	for x > 0 {
		fmt.Println(x)
		x--
	}

	for x < 10 {
		x++
		if x == 3 {
			continue // interrompe a iteração
		}
		fmt.Println(x)
		if x == 5 {
			break // interrompe o loop
		}
	}
}
