package main

import (
	"errors"
	"fmt"
)

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func main() {

	example1(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(sum(1, 2))
	fmt.Println(nameLength("Jhon"))
	fmt.Println(greeting("Jhon"))
	fmt.Println(greetingOk("Jhon"))
	fmt.Println(fibonacci(10))
	fmt.Printf("%T\n", returnParam)
	fmt.Printf("%d\n", returnFunction(1, returnParam))

	defer first() /* executed in the end of the main function */
	second()

	example2()

	/* nested functions */

	callbackFunction := func(f func()) func() {
		return f
	}

	callback(callbackFunction)(first)
}

func example1(integers ...int) {
	for _, integer := range integers {
		fmt.Printf("%d ", integer)
	}
	fmt.Println()
}

func sum(a, b int) int {
	return a + b
}

func nameLength(name string) (string, int) {
	return name, len(name)
}

func greeting(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Invalid input")
	}

	return "Heelo " + name, nil
}

func greetingOk(name string) (string, bool) {
	if len(name) == 0 {
		return "", false
	}

	return "Heelo " + name, true
}

func fibonacci(i int) int {
	if i == 0 || i == 1 {
		return i
	}

	return fibonacci(i-1) + fibonacci(i-2)
}

func byValue(i int) {
	i = i + 1
}

func byReference(i *int) {
	*i = *i + 1
}

func returnParam(i int) int {
	return i
}

func returnFunction(i int, f func(int) int) int {
	return f(i)
}

func example2() {
	/* anonymous function */
	function := func() {
		fmt.Println("anonymous function")
	}

	function()
	function = second
	function()

	func() {
		fmt.Println("anonymous function 2")
	}()
}

func callback(f func(func()) func()) func(func()) func() {
	return f
}
