package main

import "fmt"

func reciever(pipe <-chan string, done chan<- bool) {
	fmt.Println("Waiting for message")
	message := <-pipe
	fmt.Println("Got message", message)
	done <- true
}

func sender(pipe chan<- string, message string) {
	fmt.Println("Sending message", message)
	pipe <- message
	pipe <- "another string"
}

func sender1(pipe chan<- string) {
	messages := []string{"hello", "world", "this", "is", "a", "string", "slice"}
	for _, msg := range messages {
		pipe <- msg
	}
	close(pipe)
}

func reader1(pipe <-chan string, done chan<- bool) {
	for msg := range pipe {
		fmt.Println("Got : ", msg)
	}

	done <- true
}

func main() {

	pipe := make(chan string)
	done := make(chan bool)

	//go reciever(pipe, done)
	//go sender(pipe, "hello world")

	go reader1(pipe, done)
	go sender1(pipe)
	<-done
	//<-done
	fmt.Println("Exiting main function")
}
