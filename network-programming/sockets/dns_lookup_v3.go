package main

import (
	"fmt"
	"net"
)

func main() {
	args := []*struct {
		Net     string
		Address string
	}{
		{Net: "tcp4", Address: "www.debian.org:443"},
		{Net: "tcp4", Address: "www.duckduckgo.com:80"},
		{Net: "tcp4", Address: "www.archlinux.org:443"},
	}

	for i, arg := range args {
		addr, err := net.ResolveTCPAddr(arg.Net, arg.Address)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		} else {
			mask := addr.IP.DefaultMask()
			ones, bits := mask.Size()
			network := addr.IP.Mask(mask)

			fmt.Println(
				"Result ", (i + 1),
				"\n\t\tAddress is ", addr.String(),
				"\n\t\tDefault mask length is ", bits,
				"\n\t\tLeading ones count is ", ones,
				"\n\t\tMask is (hex), ", mask.String(),
				"\n\t\tNetwork is ", network.String(),
			)
		}
	}
}
