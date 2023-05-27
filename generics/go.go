package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	//Essa função vai tentar mapear as chaves do Map q for passado
	//lembrando a estrutura de um map:
	//map[TIPO]tipoValor
	for k := range m {
		//aqui ele guardou só a key no R
		r = append(r, k)
	}
	//
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	//muita loucura aqui, ele herda muitas coisas de um elemento
	//basicamente um prototype
	//ele vai ter métodos de tail para verificar oq tá na base
	//e método de head para verificar oq ta no topo
	//se a base estiver vazia, ele vai dizer que é oq vem no head
	//e o head é o element.val = v(parametro q foi recebido)
	//então oq isso quer dizer
	//q no primeiro run: ele vai adicionar um valor no head e o tial tbm vai ser esse
	//se n for o primeiro run, ele vai dizer que o lst.tail.next é igual o valor recebido
	//e o next tbm é um element , q é um ponteiro
	//isso ai é uma droga e não entendi direito esse método mas peguei a ideia

	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))
	//aqui não precisou especificar os tipos, e nem foi
	//há uma inferencia automatica nos generics

	_ = MapKeys[int, string](m)
	//Aqui nao foi preciso especificar os tipos, mas mesmo assim
	//foi especificado

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
