package main

import "fmt"

type Count func(int) int

func display(number int, count Count) {
	fmt.Println(count(number))
}

func checkCount(number int) int {
	if number > 10 {
		return number
	}
	return 0
}

func main() {
	display(11, checkCount)
}
