package main

import "fmt"

var agenda = make(map[string]int)

func main() {
	fmt.Printf("Agenda inicialmente: %v\n", agenda)

	cadastrarNaAgenda("Juninho", 51999354299)
	cadastrarNaAgenda("Pai", 51997015112)
	cadastrarNaAgenda("Pai", 51997015112)

}

func cadastrarNaAgenda(nome string, telefone int) {
	fmt.Printf("Cadastrando: %v %v\n\n", nome, telefone)
	_, cadastrado := agenda[nome]
	if cadastrado {
		fmt.Printf("O Contato %v já possuí um número cadastrado!\n\n", nome)
		return
	}
	agenda[nome] = telefone
	logAgenda()
	return
}

func logAgenda() {
	fmt.Println("Agenda foi modificada!")
	for contato, telefone := range agenda {
		fmt.Printf("Contato: %v\nTelefone: %v\n", contato, telefone)
	}
	fmt.Println()
}
