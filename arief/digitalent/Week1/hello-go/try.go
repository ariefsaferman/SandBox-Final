package main

import (
	"bufio"
	f "fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		f.Println(scanner.Text())
	}
}
