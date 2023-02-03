package main

import "fmt"

type Check func(string) bool

func display(name string, check Check) {
	if check(name) {
		fmt.Println("the name ", name, "is blocked")
	} else {
		fmt.Println("welcome", name)
	}

}

func main() {
	checkName := func(name string) bool {
		return name == "andi"
	}

	display("andi", checkName)
}
