package main

import "fmt"

func main() {
	// Simple string printing
	fmt.Println("Hello World")
	fmt.Println("1 + 1 =", 1.0+1.0)
	fmt.Println(321325*424521)
	fmt.Println(len("Hello World"))

	// Strings are indexed by 0 but return int value of byte
	fmt.Println("Hello World"[0])

	// variables: defined as var [name] [type] = [value]
	var x string
	x = "first"
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	var ( 
		a = 5
		b = 10
		c = 15 
	)
	fmt.Println(a+b+c)

	// alternatively var:=value, compiler guesses type
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

}

