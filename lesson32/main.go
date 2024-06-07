package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var clients []net.Conn

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		clients = append(clients, conn)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		fmt.Print("Received message: ", string(message))
		message = message[:len(message)-1]
			clin(message, conn)
	}
}

func clin(message string, con net.Conn) {
	// var name string
	// fmt.Scan(&name)
	for _, v := range clients {
		if v != con {
			_, err := v.Write([]byte(strings.ToUpper(message) + "FROM  SERVER\n"))
			if err != nil {
				panic(err)
				
			}
		}
	}

}

