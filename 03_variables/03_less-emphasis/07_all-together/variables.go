package main

import "fmt"

func main() {

	movieQuote := `"It's the one that says bad mother fucker."`
	statement := `Back-ticks allow double quotes, ` + movieQuote + `, and single quote apostrophe's`

	fmt.Println("movieQuote: ", movieQuote)
	fmt.Println("statement: ", statement)
}
