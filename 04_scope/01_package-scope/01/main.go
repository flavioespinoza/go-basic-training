package main

import "fmt"

var packageScope = `This is a package scope variable which is outside of func main(), so it can be used in any func`

func main() {
	fmt.Println(`Inside func Main()`, packageScope)
	foo()
}

func foo() {
	fmt.Println(`Inside func foo()`, packageScope)
}
