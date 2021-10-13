package main

import (
	"fmt"
)

var (
	n int = 19
	m int = 20
	h string = "string"
);
// global and block scope

func main() {
	var i int = 10;
	var s string;
	s = "JJLee";
	fmt.Println(i);
	fmt.Println(s);
	k := true
	fmt.Printf("%v, %T\n", k, k);
	var n int = 1000
	fmt.Println(n);
	// shadow variable n
}
