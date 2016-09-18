package main

import (
	"errors"
	"fmt"
)

func sayHello(name string) string {
	return "Hello " + name + "!"
}

func divide(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero is not possilbe")
	}
	return a / b, nil
}

func main() {
	res := sayHello("go lang")
	fmt.Println(res)

	q, r := divide(5, 2)
	fmt.Println(q, r)

	q, err := division(5, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(q)
}
