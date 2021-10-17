package main

import "fmt"

func main() {

	example1()
	example2()
	example3()
	example4()
	example5()
}

func example1() {
	fixed := [...]int{0, 1, 2}

	slice := []int{0, 1, 2}
	slice = append(slice, 3, 4)

	fmt.Printf("%T %T\n", fixed, slice)
	fmt.Printf("%d %d\n", len(fixed), len(slice))
}

func example2() {
	slice := make([]int, 5)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func example3() {
	slice := []int{0, 1, 2}

	slice = append(slice, 3, 4, 5, 6)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice
	slice = slice[0:7]

	fmt.Println(slice)

	slice = slice[0:8]

	fmt.Println(slice)

	slice = slice[0:7]

	fmt.Println(slice)

	slice = slice[0:]

	fmt.Println(slice)

	slice = slice[:7]

	fmt.Println(slice)
}

func example4() {
	slice := []int{}

	if slice == nil {
		fmt.Println("Empty")
	} else {
		fmt.Println(slice)
	}
}

func example5() {
	var slice []int

	fmt.Println(slice)

	if slice == nil {
		fmt.Println("Empty")
	} else {
		fmt.Println(slice)
	}
}
