package main

import (
	"fmt"
	"rsc.io/quote"
)

func greet() string {
	return quote.Hello()
}

func main() {
	fmt.Println(greet())
}
