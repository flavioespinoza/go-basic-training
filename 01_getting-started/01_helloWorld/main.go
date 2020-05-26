package main

import (
	"fmt"
)

// ---------------------------------------------------------
// Control Flows:
// ---------------------------------------------------------

// (1) sequence
// (2) loop: iterative
// (3) conditional

// -----------------------------------------------------------------------
// Variables: (if you declare a VARIABLE you have to use it)
// -----------------------------------------------------------------------

var format = "%T\n"

// declare a VARIABLE with IDENTIFIER "outside" and assign it a VALUE of "outside text"
var outside = `outside text`

// declare a VARIABLE with IDENTIFIER "y"
// assign it's VALUE to be 0
var y = 0

// declare a VARIABLE with IDENTIFIER "z" of TYPE "int"
// assign it's VALUE to be the ZERO_VALUE of TYPE "int"
var z int
/*
	The ZERO_VALUE is the Golang default value of a TYPE which is used if no value has been provided.

	Consider these 2 declarations:
		var y int = 0
		var z int

	The type for both is "int" and the value of both is 0:
		"y" has been explicitly assigned the numeric value 0
		"y" has been automatically assigned the type "init" because 0 is an integer
		"z"	has been explicitly assigned the type "int"
		"z" has been automatically assigned the ZERO_VALUE for type "int"

	Each ZERO_VALUE is preset in Golang by these respective TYPES:
		- false : booleans
		- 0 	: numeric types
		- "" 	: strings
		- nil 	: pointers, functions, interfaces, slices, channels, maps

	This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.

 */

// declare a VARIABLE with an inherent type "int"
var n = 7
// auto assigned TYPE "int" (can only be assigned values of TYPE "int")

// declare a VARIABLE with an inherent type "string"
var s = `Shaken, not stirred.`
// auto assigned TYPE "string" (can only be assigned values of TYPE "string")

func main() {
	// (1) sequence - enter
	fmt.Println(`sequence | enter`) //	<-- this is a statement (all fmt.Println()'s are statements)

	// --------------------------------------------------------------------------------------------------
	// SHORT OPERATOR -> ":=" (semi-colon + equal_sign) -> (declare + assign)
	// --------------------------------------------------------------------------------------------------

	// declare a VARIABLE with IDENTIFIER "inside"
	// assign it's VALUE to be a "string"
	inside := `inside text`
	fmt.Println(inside)
	fmt.Println(outside)

	// declare a VARIABLE with IDENTIFIER "m"
	// assign it's VALUE to be the math EXPRESSION: (100 + 24) / 2
	m:= (100 * 24) / 2
	fmt.Println(`m :=`, m)

	// declare a VARIABLE with IDENTIFIER "x"
	// assign it's VALUE to be an integer between 1 and 50
	x := 42
	fmt.Println(`x :=`,x)

	// assign VARIABLE "x" a new VALUE to be an integer between 51 and 99
	x = 99
	fmt.Println(`x =`,x)

	// log the VALUE of VARIABLE "y"
	fmt.Println(`var y =`, y)

	// log the VALUE of VARIABLE "z"
	fmt.Println(`var z int =>`, z)

	// log the VALUE and TYPE of VARIABLE "n"
	fmt.Print(`var n = `)
	fmt.Print(n)
	fmt.Print(`; `)
	fmt.Printf(format, n)

	// log the VALUE and TYPE of VARIABLE "s"
	fmt.Print(`var s = `)
	fmt.Print(s)
	fmt.Print(`; `)
	fmt.Printf(format, s)

	foo()

	fmt.Println(`Hello Flavio! :)`)

	// (2) loop: iterative
	for i := 0; i < 10; i++ {
		// (3) conditional
		if i%2 == 0 {
			fmt.Println(`index = `, i)
		}
	}

	logMultiple()

	bar()
	// (1) sequence | exit
}

func foo() {
	fmt.Println("I am Foo")
}

func logMultiple() {
	n, err := fmt.Println("Hello-String-Arg", 55, true)
	fmt.Println(`err -->`, err)
	fmt.Println(`n -->`, n)
}

func bar() {
	fmt.Println(`sequence | exit`)
}
