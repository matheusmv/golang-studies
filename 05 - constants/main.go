package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {

	const five int = 5

	const (
		x int = 1
		y int = 2
		w int = 3
		z int = 4
	)

	fmt.Println(x, y, w, z, five)

	const (
		a = 1
		b
		c = 3
		d
	)

	fmt.Println(a, b, c, d)

	const (
		zero int = iota
		one
		two
	)

	fmt.Println(zero, one, two)

	const (
		six int = iota + 6
		seven
		eight
	)

	fmt.Println(six, seven, eight)

	const (
		_ = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)

	fmt.Println(kb, mb, gb, tb, pb, eb)

	fmt.Println(math.Pi)
	fmt.Println(time.February)
	fmt.Println(time.Second)
	fmt.Println(time.UTC)
	fmt.Println(big.MaxExp)
	fmt.Println(big.MinExp)
}
