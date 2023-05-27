package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	val, prs := m["k2"]
	fmt.Printf("val %v, prs %v\n", val, prs)

	value, ok := m["k1"]
	fmt.Printf("value %v, ok %v\n", value, ok)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
