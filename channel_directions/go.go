package main

import "fmt"

func ping(pings chan<- string, msg string) {
	//Recebe um canal como parâmetro, só vai servir de entrada
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	//Recebe um canal carregado com informações como parâmetro e joga o que tem nele no canal de saída

	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	//Depois que carregou joga pra tela
}
