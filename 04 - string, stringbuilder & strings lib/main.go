package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "Hello World"

	fmt.Println(s)

	// length
	fmt.Println(len(s))

	fmt.Println(s[0]) // ASCII
	fmt.Printf("%c\n", s[0])

	// slice
	fmt.Println(s[0:6])
	fmt.Println(s[:6])
	fmt.Println(s[6:])

	// concatenate
	s += " Again"
	fmt.Println(s)

	// string literals
	fmt.Println("Hello \nWorld")
	fmt.Println("Hello \tWorld")
	fmt.Println("Hello \bWorld")

	// stirngs & slice of bytes
	abc := "abc"
	b := []byte(abc)

	fmt.Println(abc, b)
	fmt.Printf("%s %s\n", abc, b)

	abc2 := string(b)

	fmt.Println(abc2)

	// strings lib
	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.Compare("a", "b")) // -1 0 1
	fmt.Println(strings.HasPrefix("Hello World", "Hello"))
	fmt.Println(strings.HasSuffix("Hello World", "World"))
	fmt.Println(strings.Replace("Hello World World", "World", "Matheus", 1))
	fmt.Println(strings.Replace("Hello World World", "World", "Matheus", 2))

	stringBuilderExample()
	stringConversions()
}

func stringBuilderExample() {

	var sb strings.Builder

	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	sb.WriteString("Hello")
	sb.WriteString("World")

	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	sb.Grow(100)

	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	sb.Reset()

	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	sb.Write([]byte{65, 65, 65})
	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
}

func stringConversions() {
	// convert string to integer
	x := 123
	y := strconv.Itoa(x)

	fmt.Printf("%T %T \n", x, y)

	// convert integer to string
	w := "123"
	z, _ := strconv.Atoi(w)

	fmt.Printf("%T %T \n", w, z)
}
