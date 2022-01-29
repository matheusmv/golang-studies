package main

import (
	"fmt"
	"log"
	"net"
)

const PORT = ":8889"

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil {
		log.Fatalln(err.Error())
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(client net.Conn) {
	defer client.Close()

	var buffer [512]byte

	for {
		n, err := client.Read(buffer[:])
		if err != nil {
			return
		}

		fmt.Println(string(buffer[:]))

		_, err = client.Write(buffer[:n])
		if err != nil {
			return
		}
	}
}
