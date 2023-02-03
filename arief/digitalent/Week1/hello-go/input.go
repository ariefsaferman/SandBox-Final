package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputSample, _ := reader.ReadString('\n')

	a := strings.Split(inputSample, " ")

	fmt.Println(a)

}
