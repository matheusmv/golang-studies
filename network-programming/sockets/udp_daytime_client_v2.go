package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host:port>\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	// The Go net package recommends using these interface
	// types rather than the concrete ones

	// by using them, you lose specific methods such as SetKeepAlive
	// of TCPConn and SetReadBuffer of UDPConn, unless you do a type cast
	conn, err := net.Dial("udp", service)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = conn.Write([]byte("message"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	var buffer [512]byte

	n, err := conn.Read(buffer[:])
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(string(buffer[:n]))

	os.Exit(0)
}
