package main

import "fmt"

func modificaNomeLerdo(nome string) string {
	nome = "Lucas"
	fmt.Printf("nome parametro %v\n", &nome)
	return nome
}
func modificaNomeRapido(nome *string) {
	*nome = "Lucas"
	fmt.Printf("nome parametro %v\n", &nome)
}

func main() {
	nome := "João"
	fmt.Printf("Nome estava aqui na memória: %v\n", &nome)
	nome = modificaNomeLerdo(nome)
	// modificaNomeRapido(&nome)
	fmt.Printf("Nome Modificado: %v\n", nome)
	fmt.Printf("Ficou aqui na memória: %v\n", &nome)
}
