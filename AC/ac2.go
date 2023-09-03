package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var contatos [5]Contato
	var contato Contato
	var cmd = 1

	// lista inicial de contatos
	contatos[0] = Contato{Nome: "Victor", Email: "victor@ibmec.edu"}
	contatos[1] = Contato{Nome: "Clayton", Email: "clayton@ibmec.edu"}
	contatos[2] = Contato{Nome: "Ivo", Email: "ivo@gmail.com"}
	
	for cmd != 0 {
		fmt.Println("")
		fmt.Println("----------")
		fmt.Println("1- Adicionar contato")
		fmt.Println("2- Excluir contato")
		fmt.Println("3- Atualizar email")
		fmt.Println("4- Mostrar contatos")
		fmt.Println("#- Sair da aplicação") // # deve ser int
		fmt.Println("Informe o comando:")
		fmt.Scan(&cmd)

		switch cmd {
		case 1:
			// nome não fica gravado no struct
			leitor := bufio.NewReader(os.Stdin)
			fmt.Println("Informe o nome:")
			nome, _ := leitor.ReadString('\r')
			contatos = adicionaContato(nome, contatos)

		case 2:
			// funcionando ok
			contatos = excluiContato(contatos)

		case 3:
			// email não atualiza
			contato = escolherContato(contatos)
			contato.AlterarEmail()

		case 4:
			// funcionando ok
			for i := 0; i < 5; i++ {
				fmt.Println(i + 1,"- ", contatos[i])
			}

		default:
			fmt.Println("Você saiu da sua lista de contatos!")
			cmd = 0
		}
	}
}

type Contato struct {
	Nome  string
	Email string
}

func (c Contato) AlterarEmail() {
	fmt.Println(c.Nome)
	if c.Nome != "" {
		leitor := bufio.NewReader(os.Stdin)
		fmt.Print("Informe o novo email: ")
		email, _ := leitor.ReadString('\r')

		c.Email = email

	} else {
		fmt.Println("Contato vazio!")
	}
	
}

func escolherContato (a [5]Contato) Contato {
	var cont int
	fmt.Println("Escolha o contato (de 1 a 5): ")

	for i := 0; i < 5; i++ {
		fmt.Println(i + 1,"- ", a[i])
	}

	fmt.Scan(&cont)
	return a[cont-1]
}

func adicionaContato(nome string, a [5]Contato) [5]Contato {
	var email string
	fmt.Println(nome, "nome")
	if a[4].Nome == "" {
		for i := 0; i < 5; i++ {
			if a[i].Nome == "" {
				a[i].Nome = nome

				fmt.Println("Informe o email: ")
				fmt.Scan(&email)
				a[i].Email = email

				break
			}
		}
		fmt.Println("Contato adicionado!")
	} else {
		fmt.Println("Lista de contatos está cheia!")
	}
	return a
}

func excluiContato(a [5]Contato) [5]Contato {
	if a[0].Nome == "" {
		fmt.Println("Lista de contatos está vazia!")

	} else {
		for i := 4; i >= 0; i-- {
			if a[i].Nome != "" {
				a[i].Nome = ""
				a[i].Email = ""
				break
			}
		}
	}
	fmt.Println("Último contato válido foi excluído")
	return a
}