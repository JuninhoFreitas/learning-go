package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	//Isso aqui é uma espécie de construtor, então é um senso comum
	//sempre que usar structor usar um "construtor" para inicializar
	//os objetos
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp)
	fmt.Println(s)
}
