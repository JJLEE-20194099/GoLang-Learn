package main

import "fmt";

func main() {
	var a complex128 = 1 + 2i;
	var b = complex(1, 2);
	fmt.Printf("%v, %T\n", real(a), imag(a));
	fmt.Printf("%v, %T\n", real(b), imag(b));

	var c byte = 'H';
	fmt.Printf("%v, %T\n", c, c);
	var str string = "Hello world";
	var newStr []byte = []byte(str);
	fmt.Printf("%v, %T\n", newStr, newStr);

	var runeVar rune = 'e';
	fmt.Printf("%v, %T", runeVar, runeVar);

	


}



