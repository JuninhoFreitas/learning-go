package main

import "fmt"

func main() {
	//Nesse Buffered, tem uma diferença do padrão
	//Pois o unbuffered ('../channels' aq no projeto)
	//Ele aguarda até que tenha no código alguem disposto a receber
	//se não tiver, ele estoura um erro

	messages := make(chan string, 3)
	//mas no buffered, ao adicionar esse segundo parametro no make
	//ele se torna um buffered
	//e já não tem necessidade de que alguem receba essa mensagem

	messages <- "buffered"
	messages <- "channel"
	messages <- "wow"

	fmt.Println(<-messages)
	// fmt.Println(<-messages)

	func() { fmt.Println(<-messages) }()
}
