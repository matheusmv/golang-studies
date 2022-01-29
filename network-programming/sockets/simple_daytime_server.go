package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const PORT = ":8888"

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", PORT)
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

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
