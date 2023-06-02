package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		//[7, 42]
		//7 + 3 dá 10, então primeiro resultado
		// f1 worked: 10
		if r, e := f1(i); e != nil {
			// f1 failed: can't work with 42
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		//[7, 42]
		//7 + 3 dá 10, então primeiro resultado
		// f1 worked: 10
		if r, e := f2(i); e != nil {
			//42 não pode
			//Então deu o erro implementado
			// custom error deve ter o método Error() implementado
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
