package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	myIpAddress    = "0.0.0.0"
	ipv4HeaderSize = 20
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host>", os.Args[0])
		os.Exit(1)
	}

	localAddr, err := net.ResolveIPAddr("ip4", myIpAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}

	remoteAddr, err := net.ResolveIPAddr("ip4", os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
	}

	conn, err := net.DialIP("ip4:icmp", localAddr, remoteAddr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer conn.Close()

	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum
	msg[3] = 0  // checksum
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1] (arbitrary)
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1] (arbitrary)
	len := 8

	check := checkSum(msg[:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[:len])
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Print("Message sent: ")
	for n := 0; n < 8; n++ {
		fmt.Print(" ", msg[n])
	}
	fmt.Println()

	size, err := conn.Read(msg[:])
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Print("Message received: ")
	for n := ipv4HeaderSize; n < size; n++ {
		fmt.Print(" ", msg[n])
	}
	fmt.Println()

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 0; n < len(msg); n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)

	var answer uint16 = uint16(^sum)
	return answer
}
