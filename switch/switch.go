package main

import "fmt"

func main() {

	exit := false

	for !exit {
		fmt.Println("Enter your choice")
		var choice string
		fmt.Scanf("%s", &choice)

		switch choice {
		case "yes", "Yes":
			fmt.Println("User entered yes")
		case "no", "No":
			fmt.Println("User entered no")
		case "exit":
			fmt.Println("bye")
			exit = true
		default:
			fmt.Println("unknown value")
		}
	}
}
