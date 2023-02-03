package main

import (
	"bufio"
	"fmt"
	"os"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func menu() {
	fmt.Print("Activity Reporter\n" +
		"1. Setup\n" +
		"2. Action\n" +
		"3. Display\n" +
		"4. Trending\n" +
		"5. Exit\n")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	exit := false

	for !exit {
		menu()
		input := promptInput(scanner, "input menu: ")

		switch input {
		case "1":
		case "2":
		case "3":
		case "4":
		case "5":
		default:
			fmt.Println("invalid menu")
		}
	}
}
