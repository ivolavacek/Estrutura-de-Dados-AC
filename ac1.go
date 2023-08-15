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
	fmt.Println("Digite o número que corresponda a um dia da semana")
	fmt.Scan(&z)
	fmt.Println(diaSemana(z))

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
}

func diaSemana(dia int) string{
	if dia == 1 {
		return "Domingo"
	} else if dia == 2 {
		return "Segunda-feira"
	} else if dia == 3 {
		return "Terça-feira"
	} else if dia == 4 {
		return "Quarta-feira"
	} else if dia == 5 {
		return "Quinta-feira"
	} else if dia == 6 {
		return "Sexta-feira"
	} else if dia == 7 {
		return "Sábado"
	} else {
		return "Valor Inválido!"
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