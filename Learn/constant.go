package main

import (
	"fmt"
)

const (
	a = 42
	b = 54
)

const (
	red = iota
	yellow
	black
	white
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

	// const error = math.sqrt(4)
	// fmt.Printf("%v, %T\n", error, error)
	// go lang là ngôn ngữ biên dịch khi biên dịch thì hàm sqrt.math() chưa được chạy 
	// nên gán cho const thì sẽ sinh ra lỗi

	fmt.Printf("%v, %T\n", white, white)

	
}	
