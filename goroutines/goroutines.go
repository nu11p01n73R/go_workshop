package main

import (
	"fmt"
	"time"
)

func say(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str, " = ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	say("main function")

	go say("say call one")
	go say("say call two")

	go func(str string) {
		for i := 0; i < 3; i++ {
			fmt.Println(str, " = ", i)
			time.Sleep(time.Second)
		}
	}("anonymous call one")

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}
