package main

import "fmt"

func calcMax(x int) int {
	return 42 + x
}

func main() {
	max := calcMax(7)
	fmt.Println(max) // max is now the result, not the function
}

// don't do this; bad coding practice to shadow variables
