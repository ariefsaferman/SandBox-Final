package main

import "fmt"

const MinValue = 10

func main() {
	fmt.Println(MinValue)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
