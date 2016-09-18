package main

import "fmt"

func main() {

	var i int = 1

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 10; j > 0; j-- {
		fmt.Println(j)
	}

}
