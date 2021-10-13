package main

import "fmt"
import "strconv"

var (
	n int = 19
	m int = 20
	h string = "string"
);
// global and block scope

var N int = 10

func main() {
	// var i int = 10;
	// var s string;
	// s = "JJLee";
	// fmt.Println(i);
	// fmt.Println(s);
	// k := true
	// fmt.Printf("%v, %T\n", k, k);
	// var n int = 1000
	// fmt.Println(n);
	// shadow variable n

	var i int = 1000;
	// var j float32 = float32(i);
	var str = strconv.Itoa(i);
	// var t = string(i);
	fmt.Printf("%v, %T\n", str, str);
	// fmt.Printf("%v, %T\n", t, t);


}
