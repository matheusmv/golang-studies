package main

import "fmt"

func main() {

	example1()
	example2()
	example3()
	example4()
}

func example1() {
	var m map[string]string

	fmt.Println(m == nil)

	m = map[string]string{}

	fmt.Println(m == nil)
	fmt.Println(m)
	fmt.Println(len(m))

	m = make(map[string]string, 5)

	fmt.Println(m == nil)
	fmt.Println(m)
	fmt.Println(len(m))
}

func example2() {
	m := map[string]string{
		"Wallace": "Programmer",
		"Jhon":    "Admin",
	}

	fmt.Println(m == nil)
	fmt.Println(m)
	fmt.Println(len(m))

	m["Alex"] = "Intern"

	fmt.Println(m)
	fmt.Println(len(m))

	m["Alex"] = "Junior Developer"

	fmt.Println(m)
	fmt.Println(len(m))

	delete(m, "Wallace")

	fmt.Println(m)
	fmt.Println(len(m))

	m["Alex"] += " Content Creator"

	fmt.Println(m)
	fmt.Println(len(m))
}

func example3() {
	m := map[string]string{
		"Wallace": "Programmer",
		"Jhon":    "Admin",
		"Alex":    "Junior Developer",
	}

	for key, value := range m {
		fmt.Printf("%s -> %s\n", key, value)
	}
}

func example4() {
	m := map[string]string{
		"Wallace": "Programmer",
		"Jhon":    "Admin",
		"Alex":    "Junior Developer",
	}

	value, isPresent := m["Johnny"]

	if isPresent {
		fmt.Println(value)
	} else {
		fmt.Println("key dont exist")
	}
}
