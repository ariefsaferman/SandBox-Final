package main

import "fmt"

func main() {

	fmt.Println(factorial(5))
}

func factorial(angka int) int {
	if angka == 1 {
		return 1
	} else {
		return angka * factorial(angka-1)
	}
}
