package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) (int, string) {
	return (a + b + c), ("qual foi")
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res, str := plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	fmt.Println("STR", str)
}
