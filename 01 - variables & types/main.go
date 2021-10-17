package main

import (
	"fmt"
)

func fibonacciReq(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return (fibonacciReq(n-1) + fibonacciReq(n-2))
	}
}

func fibonacciSeq(max int) {
	var f1, f2 = 0, 1

	for i := 0; i < max; i++ {
		fmt.Print(f1, " ")
		f1, f2 = f2, (f1 + f2)
	}
}

func main() {

	// var name type = expression
	var age, name = 25, "Jhon"

	fmt.Println(name, age)

	// name := expression
	boolean := false

	fmt.Println(boolean)

	// pointers
	x := 1
	p1 := &x

	fmt.Println(p1)
	fmt.Println(*p1)

	var y int = 1
	var p2 *int = &y

	fmt.Println(p2)
	fmt.Println(*p2)

	// type
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)

	fmt.Println(c)
}
