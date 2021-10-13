package main

import "fmt";

func main() {
	var a complex128 = 1 + 2i;
	var b = complex(1, 2);
	fmt.Printf("%v, %T\n", real(a), imag(a));
	fmt.Printf("%v, %T\n", real(b), imag(b));

}



