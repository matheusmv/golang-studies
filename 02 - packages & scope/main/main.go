package main

import (
	"fmt"

	"matheusmv.com/packages-modules/integers"
)

var five = 5

func main() {

	{
		var integer = 2

		fmt.Println(integer)

		fmt.Println(five)
	}

	var integer = 1

	fmt.Println(integer)

	fmt.Println(five)

	fmt.Println(integers.Three)

	foo()
}

func foo() {
	fmt.Println(five)
}
