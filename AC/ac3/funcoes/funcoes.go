package funcoes

import (
	"fmt"
	c "ac3/contato"
)

// Função criada para receber o endereço de um contato existente para, em outra função, mudar o email
func EscolherContato (endContato *[5]c.Contato) *c.Contato{
	var cont int

	fmt.Print("Escolha o contato (de 1 a 5): ")

	for i := 0; i < 5; i++ {
		fmt.Println(i + 1,"- ", endContato[i])
	}

	fmt.Scan(&cont)
	return &endContato[cont-1]
}

// Confere se há espaço no array, se tiver adiciona um novo contato, com campos de 'nome' e 'email'
func AdicionaContato(nome string, endContato *[5]c.Contato) {
	var email string

	if endContato[4].Nome == "" {
		for i := 0; i < 5; i++ {
			if endContato[i].Nome == "" {
				endContato[i].Nome = nome

				fmt.Print("Informe o email: ")
				fmt.Scan(&email)
				endContato[i].Email = email

				break
			}
		}
		fmt.Println("Contato adicionado!")
	} else {
		fmt.Println("Lista de contatos está cheia!")
	}
}

// Confere se o array está vazio, se não, apaga último contato
func ExcluiContato(endContato *[5]c.Contato) {
	if endContato[0].Nome == "" {
		fmt.Println("Lista de contatos está vazia!")

	} else {
		for i := 4; i >= 0; i-- {
			if endContato[i].Nome != "" {
				endContato[i].Nome = ""
				endContato[i].Email = ""
				break
			}
		}
	}
	fmt.Println("Último contato válido foi excluído")
}