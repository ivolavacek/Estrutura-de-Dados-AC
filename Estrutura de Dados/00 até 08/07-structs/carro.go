package main

import "fmt"

type Carro struct {
	Chassi     string
	Cor        string
	Potencia   int
	Ano        int
	EstaLigado bool
	Velocidade int
}

func (c *Carro) Ligar() {
	c.EstaLigado = true
	fmt.Println("Ligou o carro")
}

func (c *Carro) Acelerar(valor int) {
	if c.EstaLigado {
		c.Velocidade += valor
	} else {
		fmt.Println("Liga o carro, Zé Mané")
	}
}

func (c *Carro) Desligar() {
	c.EstaLigado = false
	c.Velocidade = 0
	fmt.Println("Parou o carro...")
}

func main() {
	c := Carro{
		Chassi:     "12234",
		Cor:        "preto",
		Potencia:   200,
		Ano:        2023,
		EstaLigado: false,
		Velocidade: 0,
	}

	fmt.Println(c)
	c.Acelerar(15)
	fmt.Println(c)
	c.Ligar()
	fmt.Println(c)

	c.Acelerar(15)
	fmt.Println(c)

	c.Desligar()
}
