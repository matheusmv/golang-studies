package main

import (
	"bytes"
	"fmt"
	"io"
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

	// The Go net package recommends using these interface
	// types rather than the concrete ones

	// by using them, you lose specific methods such as SetKeepAlive
	// of TCPConn and SetReadBuffer of UDPConn, unless you do a type cast
	conn, err := net.Dial("tcp", service)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer conn.Close()

	_, err = conn.Write([]byte(HEAD_REQUEST))
	if err != nil {
		log.Fatalln(err.Error())
	}

	result, err := readFully(conn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(string(result))

	os.Exit(0)
}

func readFully(conn net.Conn) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	var buffer [512]byte

	for {
		n, err := conn.Read(buffer[:])
		result.Write(buffer[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}
