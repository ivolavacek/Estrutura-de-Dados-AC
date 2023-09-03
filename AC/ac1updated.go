package main

import "fmt"

func main(){
	// Questão do número primo
	var x int
	fmt.Println("Informe o número que deseja saber se é primo: ")
	fmt.Scan(&x)
	fmt.Println(ePrimo(x))

	// Questão do número de Fibonacci
	var y int
	fmt.Println("Qual 'n-ésimo' número do número de Fibonacci que deseja saber: ")
	fmt.Scan(&y)
	fmt.Println(fibo(y))

	//Questão do dia da semana
	var z int
	fmt.Println("Digite o número que corresponda a um dia da semana: ")
	fmt.Scan(&z)
	diaSemana(z)

	// Questão do ano bissexto
	var ano int
	fmt.Println("Digite o ano que deseja saber se é bissexto: ")
	fmt.Scan(&ano)
	fmt.Println(eBissexto(ano))
}

func ePrimo(num int) bool{
	var divisores []int

	for i := 1; i <= num; i++{
		if num % i == 0{
			divisores = append(divisores, i)
		}
	}

	if len(divisores) == 2 {
		return true
	} else {
		fmt.Println(divisores)
		return false
	}

	/*
	var primo = true

	for i := 2; i < num; i++ {
		if num % i == 0 {
			primo = false
			fmt.Print(i, " ")
		}
	}

	fmt.Println("")
	return primo
	*/
}

func fibo(num int) int{
	fibonacci := []int{1,1,2,3}
	var proxNum int

	if num >= 5 {
		for i := 5; i <= num; i++{
			proxNum = fibonacci[i-2] + fibonacci[i-3]
			fibonacci = append(fibonacci, proxNum)
		}
	}

	return fibonacci[num-1]

	/*
	var resultado, ult, penult int
	ult, penult = 1, 1

	for i := 3; i <= num; i++ {
		resultado = ult + penult
		ult, penult = resultado, ult
	}

	return resultado
	*/
}

func diaSemana(dia int) {
	switch dia {
	case 1:
		fmt.Println("1-", "Domingo")
	case 2:
		fmt.Println("2-", "Segunda-feira")
	case 3:
		fmt.Println("3-", "Terça-feira")
	case 4:
		fmt.Println("4-", "Quarta-feira")
	case 5:
		fmt.Println("5-", "Quinta-feira")
	case 6:
		fmt.Println("6-", "Sexta-feira")
	case 7:
		fmt.Println("7-", "Sábado")
	default:
		fmt.Println("Valor inválido!")
	}
}

func eBissexto(ano int) bool {
	var bissexto = false

	if ano % 4 == 0 {
		bissexto = true
	}
	if ano % 100 == 0 && ano % 400 != 0 {
		bissexto = false
	}

	return bissexto
}