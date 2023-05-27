package main

import "fmt"

type tipo int

var x tipo

func main() {
	fmt.Printf("%v %T \n", x, x)
	x = 42
	fmt.Println(x)
}
