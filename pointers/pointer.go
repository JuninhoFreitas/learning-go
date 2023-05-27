package main

import "fmt"

func zeroval(ival int) {
	ival = 0
	fmt.Println(&ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	fmt.Println("Pointed:", &i)
	//i tá em um determinado endereço, mas dentro de zeroval, é alocado em outro lugar
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	//como o zeroptr aponta para o endereço que tá alocado o i
	//ele acaba modificando o i que era 1 para 0
	//Pois é um ponteiro
	//Economia de memória, pois não precisamos fazer igual no js
	// ex:
	// nome = "João"
	// nome = () => {let newName = "Junior"; return newName}
	// para modificar o nome, podemos só apontar
	fmt.Println("pointer:", &i)
}
