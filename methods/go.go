package main

import "fmt"

type rect struct {
	width, height int
}

//tanto faz se for ponteiro ou não, ele faz conversão
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("R: ", r)
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("RP: ", rp)
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	rp.height = 100
	rp.width = 100
	// fmt.Println("RP: ", rp)
	// fmt.Println("area: ", rp.area())
	// fmt.Println("perim:", rp.perim())

	// r é semelhante a um objeto com funções dentro
	// r.width, r.height, r.area, r.perim
	fmt.Println("R: ", r)
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
}
