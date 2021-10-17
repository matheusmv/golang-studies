package main

import (
	"fmt"
	"math"
)

func main() {

	intergers()
	comparatorOperators()
	logicalOperators()
	arithmeticOperators()
	binaryOperators()
	floatingPointNumbers()
	mathPackage()
	complexNumbers()
}

func intergers() {
	fmt.Println("integers")

	var x int = 1 // int32 int64
	var y = 1
	z := 1

	fmt.Println(x, y, z)

	var x8 int8 = math.MaxInt8
	var x16 int16 = math.MaxInt16
	var x32 int32 = math.MaxInt32
	var x64 int64 = math.MaxInt64

	fmt.Println("max 8 bits int = ", x8)
	fmt.Println("max 16 bits int = ", x16)
	fmt.Println("max 32 bits int = ", x32)
	fmt.Println("max 64 bits int = ", x64)

	var nx8 int8 = math.MinInt8
	var nx16 int16 = math.MinInt16
	var nx32 int32 = math.MinInt32
	var nx64 int64 = math.MinInt64

	fmt.Println("min 8 bits int = ", nx8)
	fmt.Println("min 16 bits int = ", nx16)
	fmt.Println("min 32 bits int = ", nx32)
	fmt.Println("min 64 bits int = ", nx64)

	var ux8 uint8 = math.MaxUint8
	var ux16 uint16 = math.MaxUint16
	var ux32 uint32 = math.MaxUint32
	var ux64 uint64 = math.MaxUint64

	fmt.Println("min unsigned 8 bits int = ", ux8)
	fmt.Println("min unsigned 16 bits int = ", ux16)
	fmt.Println("min unsigned 32 bits int = ", ux32)
	fmt.Println("min unsigned 64 bits int = ", ux64)
}

func comparatorOperators() {
	fmt.Println("comparatorOperators")

	x, y := 1, 2

	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)
}

func logicalOperators() {
	fmt.Println("logicalOperators")

	x, y := true, false

	fmt.Println(x && y)
	fmt.Println(x || y)
	fmt.Println(!x)
}

func arithmeticOperators() {
	fmt.Println("arithmeticOperators")

	x, y := 1, 2

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
}

func binaryOperators() {
	fmt.Println("binaryOperators")

	x, y := 1, 2

	fmt.Println(x & y)  // bitwise AND
	fmt.Println(x | y)  // bitwise OR
	fmt.Println(x ^ y)  // bitwise XOR
	fmt.Println(x &^ y) // bit clear (AND NOT)
	fmt.Println(x << y) // left shift
	fmt.Println(x >> y) // right shit
}

func floatingPointNumbers() {
	fmt.Println("floatingPointNumbers")

	phi := math.Phi
	pi := math.Pi

	fmt.Printf("%T\n", phi)
	fmt.Printf("%T\n", pi)

	var f32 float32 = math.MaxFloat32
	var f64 float64 = math.MaxFloat64

	fmt.Println(f32)
	fmt.Println(f64)
}

func mathPackage() {
	fmt.Println("mathPackage")

	fmt.Println(math.Ceil(3.56874))
	fmt.Println(math.Floor(3.56874))
	fmt.Println(math.Min(3, 0))
	fmt.Println(math.Max(3, 0))
	fmt.Println(math.Abs(-3))
	fmt.Println(math.Sqrt(100))
	fmt.Println(math.Pow(2, 3))
}

func complexNumbers() {
	fmt.Println("complexNumbers")

	fmt.Println(complex(1, 2))
}
