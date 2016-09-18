package main

import (
	"fmt"
	"net"
)

func connectionHandler(conn net.Conn) {
	buff := make([]byte, 1024)

	len, err := conn.Read(buff)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read %d bytes from client\n", len)
	conn.Write([]byte("Read data : " + string(buff)))
	conn.Close()
}

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	defer listen.Close()
	fmt.Println("Listening for connections")
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Accepting new connection")
		go connectionHandler(conn)
	}
}
