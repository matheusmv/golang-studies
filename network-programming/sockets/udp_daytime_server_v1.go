package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const PORT = ":8884"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		log.Fatalln(err.Error())
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		handleClient(conn)
	}
}

func handleClient(client *net.UDPConn) {
	var buffer [512]byte

	_, addr, err := client.ReadFromUDP(buffer[:])
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	daytime := time.Now().String()

	_, err = client.WriteToUDP([]byte(daytime), addr)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
}
