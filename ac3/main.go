package main

import (
	"fmt"
	c "ac3/contato"
	f "ac3/funcoes"
)

func main() {
	var contatos [5]c.Contato
	var cmd = 1

	// lista inicial de contatos
	contatos[0] = c.Contato{Nome: "Victor", Email: "victor@ibmec.edu"}
	contatos[1] = c.Contato{Nome: "Clayton", Email: "clayton@ibmec.edu"}
	contatos[2] = c.Contato{Nome: "Ivo", Email: "ivo@gmail.com"}

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
			//Modo novo de usar fmt.Scan - NÃO ESTÁ FUNCIONANDO
			var nome string
			leitor := bufio.NewReader(os.Stdin)
			fmt.Print("Informe o nome:")
			nome, _ = leitor.ReadString('\r')
			*/
			var nome string
			fmt.Print("Contato de: ");
			fmt.Scan(&nome)
			f.AdicionaContato(nome, &contatos)

		case 2:
			// Exclui contato - funcionando ok
			f.ExcluiContato(&contatos)

		case 3:
			// Atualiza email de um contato específico - funcionando ok
			contato := f.EscolherContato(&contatos)
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