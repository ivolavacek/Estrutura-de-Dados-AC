package main

import (
	"fmt"
	"math"
)

type Pessoa struct {
	Nome	string // isso é um campo
	Idade	int
	Altura	float64
}

func (p Pessoa) AvancarIdade() {
	p.Idade++
}

type Ponto struct {
	X, Y	int
}

type Retangulo struct {
	Ponto //campo anônimo
	Largura, Altura	int
}

type Circulo struct {
	Raio	float64
}

func ( c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	circ := Circulo{
		Raio:	1.65,
	}
	fmt.Println(circ.Area())

	p := Pessoa {
		Nome:	"Alice",
		Idade:	25,
		Altura:	1.65,
	}

	fmt.Println(p)
	fmt.Println(p.Nome)
	fmt.Println(p.Altura)

	p.AvancarIdade()
	fmt.Println(p)

	r := Retangulo {
		Ponto:	Ponto{X: 1, Y: 2}, // ponto anônimo
		Largura:	60,
		Altura:	30,
	}
	fmt.Println(r)
}