package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const HEAD_REQUEST = "HEAD / HTTP/1.0\r\n\r\n"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatalln(err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer conn.Close()

	_, err = conn.Write([]byte(HEAD_REQUEST))
	if err != nil {
		log.Fatalln(err.Error())
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(string(result))

	os.Exit(0)
}
