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

	udpAddr, err := net.ResolveUDPAddr("udp", service)
	if err != nil {
		log.Fatalln(err.Error())
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
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
