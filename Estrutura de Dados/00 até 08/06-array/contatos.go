package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome string
}

func main() {
	var pessoas [3]Pessoa

	pessoas[0] = Pessoa{Nome: "Victor"}

	if (pessoas[0] == Pessoa{}) {
		fmt.Println("Elemento 0 está vazio!")
	} else {
		fmt.Println("Elemento 0 não está vazio!")
	}

	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Informe o seu nome: ")
	msg, _ := leitor.ReadString('\n') // '\n' é um byte
	fmt.Println(msg)
}

func adicionaPessoa(nome string, a [3]Pessoa) [3]Pessoa {
	a[0] = Pessoa{Nome: nome}
	return a
}
