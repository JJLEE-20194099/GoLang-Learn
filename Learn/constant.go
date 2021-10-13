package main

import "fmt"

const (
	a = 42
	b = 54
)

func main() {
	const i int16 = 42
	const j = "Hello"
	const a = 21
	// shadow constant
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

}
