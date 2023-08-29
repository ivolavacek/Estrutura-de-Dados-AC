package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
	var contatos [5]Contato
	var cmd = 1

	// lista inicial de contatos
	contatos[0] = Contato{Nome: "Victor", Email: "victor@ibmec.edu"}
	contatos[1] = Contato{Nome: "Clayton", Email: "clayton@ibmec.edu"}
	contatos[2] = Contato{Nome: "Ivo", Email: "ivo@gmail.com"}
	
	for cmd != 0 {
		fmt.Println("")
		fmt.Println("----------")
		fmt.Println("")
		fmt.Println("1- Adicionar contato")
		fmt.Println("2- Excluir contato")
		fmt.Println("3- Atualizar email")
		fmt.Println("4- Mostrar contatos")
		fmt.Println("#- Sair da aplicação") // # deve ser int
		fmt.Println("")
		fmt.Print("Informe o comando: ")
		fmt.Scan(&cmd)

		switch cmd {
		case 1:
			// Adiciona contato - funcionando ok
			/*
			Modo novo de usar fmt.Scan - NÃO ESTÁ FUNCIONANDO
			leitor := bufio.NewReader(os.Stdin)
			fmt.Println("Informe o nome:")
			nome, _ := leitor.ReadString('\r')
			*/
			var nome string

			fmt.Print("Contato de: ");
			fmt.Scan(&nome)
			adicionaContato(nome, &contatos)

		case 2:
			// Exclui contato - funcionando ok
			excluiContato(&contatos)

		case 3:
			// Atualiza email de um contato específico - funcionando ok
			contato := escolherContato(&contatos)
			contato.AlterarEmail()

		case 4:
			// Mostra lista de contatos - funcionando ok
			for i := 0; i < 5; i++ {
				fmt.Println(i + 1,"- ", contatos[i])
			}

		default:
			fmt.Println("Você saiu da sua lista de contatos!")
			cmd = 0
		}
	}
}

// Struct do 'Contato' com 2 campos
type Contato struct {
	Nome  string
	Email string
}

// Função que, a partir do endereço de um contato, muda o seu email
func (c *Contato) AlterarEmail() { // *Contato -> mudança sugerida pelo professor
	var email string
	
	if c.Nome != "" {
		fmt.Print("Informe o novo email: ")
		fmt.Scan(&email)

		c.Email = email

	} else {
		fmt.Println("Contato vazio!")
	}
}

// Função criada para receber o endereço de um contato existente para, em outra função, mudar o email
func escolherContato (endContato *[5]Contato) *Contato{
	var cont int

	fmt.Print("Escolha o contato (de 1 a 5): ")

	for i := 0; i < 5; i++ {
		fmt.Println(i + 1,"- ", endContato[i])
	}

	fmt.Scan(&cont)
	return &endContato[cont-1]
}

// Confere se há espaço no array, se tiver adiciona um novo contato, com campos de 'nome' e 'email'
func adicionaContato(nome string, endContato *[5]Contato) {
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
func excluiContato(endContato *[5]Contato) {
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