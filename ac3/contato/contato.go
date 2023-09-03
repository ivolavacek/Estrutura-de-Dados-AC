package contato

import "fmt"

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