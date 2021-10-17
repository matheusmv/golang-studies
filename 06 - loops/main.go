package main

import "fmt"

func main() {

	// for loops
	example1()
	example2()
	example3()
	example4()
	example5(10)
	example6()
	example7()
	example8()
	example9()

	// foreach loops
	exampleForeach1()
	exampleForeach2()
	exampleForeach3()
}

func example1() {
	i := 0
	for {
		fmt.Printf("%d ", i)
		i += 1

		if i == 10 {
			break
		}
	}
	fmt.Println()
}

func example2() {
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i += 1
	}
	fmt.Println()
}

func example3() {
	for i := 0; i < 10; {
		fmt.Printf("%d ", i)
		i += 1
	}
	fmt.Println()
}

func example4() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func example5(n int) {
	for i, j, total := 0, 1, 0; total < n; i, j, total = j, (i + j), (total + 1) {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func example6() {
	s := "Hello World"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %c \t", i, s[i])
	}
	fmt.Println()
}

func example7() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}

func example8() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break
			}

			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}

func example9() {
iForLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break iForLoop
			}

			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}

func exampleForeach1() {
	s := "Hello World"

	for i, c := range s {
		fmt.Printf("%d %c \t", i, c)
	}
	fmt.Println()
}

func exampleForeach2() {
	s := "Hello World"

	for i, c := range s {
		if c == ' ' {
			break
		}

		fmt.Printf("%d %c \t", i, c)
	}
	fmt.Println()
}

func exampleForeach3() {
	s := "Hello World"

	for i, c := range s {
		if c == ' ' {
			continue
		}

		fmt.Printf("%d %c \t", i, c)
	}
	fmt.Println()
}
