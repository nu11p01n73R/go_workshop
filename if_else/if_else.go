package main

import "fmt"

func main() {

	var num int = 7

	if num%2 == 0 {
		fmt.Println(num, " is even.")
	} else {
		fmt.Println(num, " is odd.")
	}

	var input int

	fmt.Println("Enter a number")
	fmt.Scanf("%d", &input)
	if input%4 == 0 {
		fmt.Printf("%d is dividable by 4\n", input)
	}
}
