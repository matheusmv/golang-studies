package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const PORT = ":8884"

func main() {
	// The Go net package recommends using these interface
	// types rather than the concrete ones

	// by using them, you lose specific methods such as SetKeepAlive
	// of TCPConn and SetReadBuffer of UDPConn, unless you do a type cast
	conn, err := net.ListenPacket("udp", PORT)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		handleClient(conn)
	}
}

func handleClient(client net.PacketConn) {
	var buffer [512]byte

	_, addr, err := client.ReadFrom(buffer[:])
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	daytime := time.Now().String()

	_, err = client.WriteTo([]byte(daytime), addr)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
}
