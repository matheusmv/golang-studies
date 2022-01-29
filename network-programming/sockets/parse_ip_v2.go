package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <ip> <ones> <bits>\n", os.Args[0])
		os.Exit(1)
	}

	ip := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])

	addr := net.ParseIP(ip)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	mask := net.CIDRMask(ones, bits)
	network := addr.Mask(mask)

	fmt.Println(
		"Address is ", addr.String(),
		"\nMask length is ", bits,
		"\nLeading ones count is ", ones,
		"\nMask is (hex), ", mask.String(),
		"\nNetwork is ", network.String(),
	)

	os.Exit(0)
}
