package main

import "fmt"

type tipo int

var x tipo
var y int

func main() {
	fmt.Printf("%v %T \n", x, x)
	x = 42
	fmt.Printf("%v %T \n", x, x)

	y = int(x)
	fmt.Printf("%v %T \n", y, y)
}
