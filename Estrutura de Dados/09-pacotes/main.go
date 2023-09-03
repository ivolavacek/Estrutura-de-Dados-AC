package main

import (
	f "fmt"
	"projeto/utils"
	"projeto/utils/outroUtils"
)

func main() {
	f.Println("Olá, mundo!")

	f.Println(utils.Somar(4.5, 2.2))
	f.Println(utils.Multiplicar(4.5, 2.2))
	f.Println(outroUtils.Dividir(8.2, 4.1))
}