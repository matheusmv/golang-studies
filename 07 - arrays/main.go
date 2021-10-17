package main

import "fmt"

func main() {

	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
}

func example1() {
	var a [3]int

	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(len(a))

	a[0] = 1

	fmt.Println(a)
}

func example2() {
	var a [3]int
	aCopy := a // arrays are values in Golang

	fmt.Println(a == aCopy)
	fmt.Println(&a == &aCopy)

	fmt.Println(a, aCopy)

	a[0] = 2

	fmt.Println(a, aCopy)

	aCopy[0] = 4

	fmt.Println(a, aCopy)
}

func example3() {
	b := [4]int{1, 2, 3, 4}

	fmt.Println(b)

	c := [...]int{2, 4, 6, 8, 10}

	fmt.Println(c)

	d := [3]int{1, 2}

	fmt.Println(d)
}

func example4() {
	a := [...]int{2, 4, 6, 8, 10}

	for i, v := range a {
		fmt.Println(i, "->", v)
	}
}

func example5() {
	array1 := [...]int{3: 1}

	fmt.Println(array1)

	array2 := [...]int{1: 1, 4: 5}

	fmt.Println(array2)
}

func example6() {
	strArray := [...]string{"string1", "stirng2"}

	fmt.Println(strArray)
}

func example7() {

	array2d := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(array2d)

	array3d := [3][2][2]int{array2d, array2d, array2d}

	fmt.Println(array3d)
}
