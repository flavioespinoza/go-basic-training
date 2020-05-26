package main

import (
	"fmt"
)

// types
type hotdog int

// set VARIABLES "a" and "b" to numeric ZERO_VALUES
var a int
var b hotdog

func main() {

	// print initial values
	fmt.Println(`var a =`, a, `; type = int`)
	fmt.Println(`var b =`, b, `; type = hotdog`)

	// change VALUE of "b" to an integer > 0
	b = 25
	fmt.Println(`b =`, b)

	// ------------------------------------------------------------------------------------------------------------
	// Example Error from trying to assign a VALUE with an incompatable TYPE
	// ------------------------------------------------------------------------------------------------------------
	/*
		Consider this assignment

			a = b

		This throws an error because VARIABLE "a" has TYPE "int" and VARIABLE "b" has the custom TYPE "hotdog"

			[error] cannot use var "b" (type "hotdog") as type "int" in assignment

		However, since TYPE "hotdog" has the UNDERLYING_TYPE "int" we can use a "CONVERSION" to get
		the desired result.

	 */
	// ------------------------------------------------------------------------------------------------------------
	// Conversion
	// ------------------------------------------------------------------------------------------------------------
	/*
		Overview:
			VARIABLE "a" has TYPE "int"
			VARIABLE "b" has TYPE "hotdog" which has UNDERLYING_TYPE "int"

		Strategy:
			We can use TYPE "int" as a function and pass in "b" to evaluate and verify that it does in-fact
			have the UNDERLYING_TYPE of "int". If this is the case the "init" function will CONVERT "b" to use
			the UNDERLYING_TYPE of "int" and assign the VALUE of "b" to the VARIABLE "a"

	 */
	// ------------------------------------------------------------------------------------------------------------
	a = int(b)
	fmt.Println(`[Conversion]: a = int(b)`)
	fmt.Println(`[Result]: a =`, a)

}

